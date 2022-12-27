
// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "./Math.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/security/Pausable.sol";
// import "hardhat/console.sol";

contract Core is Ownable, Pausable, Math {
    struct CDP {
        address owner;
        uint256 collateral;
        uint256 debt;
    }

    mapping(address => CDP) cdps;

    uint256 public totalCollateral;
    uint256 public totalDebt;

    uint256 public maxDebt;  // for whole contract
    uint256 public maxDebtCDP;  // for each CDP
    
    uint256 public liquidateFee;    // RAY
    uint256 public collateralRate;  // RAY
    uint256 public collateralPrice; // RAY

    event CDPDeleted(address indexed user);
    event CDPFunded(address indexed user, uint256 collateral, uint256 debt);
    event CDPDefunded(address indexed user, uint256 collateral, uint256 debt);
    event CDPLiquidated(
        address indexed liquidator, 
        address indexed user, 
        uint256 liqdebt, 
        uint256 liquidatorCashback,
        uint256 liquidatorMargin,
        uint256 debtorCashback
    );

    constructor(uint256 md, uint256 mdc, uint256 r, uint256 p, uint256 f) {
        maxDebt = md;
        maxDebtCDP = mdc;
        setCollateralRate(r);
        setCollateralPrice(p);
        setLiquidateFee(f);
    }

    function setCollateralRate(uint256 rate) public onlyOwner {
        collateralRate = rate;
    }

    function setCollateralPrice(uint256 price) public onlyOwner {
        collateralPrice = price;
    }

    function setLiquidateFee(uint256 fee) public onlyOwner {
        liquidateFee = fee;
    }

    function getBalance(address user) public view returns(uint256, uint256) {
        CDP storage cdp = cdps[user];
        return (cdp.collateral, cdp.debt);
    }
 
    function fund(address user, uint256 collateral, uint256 debt) 
        external 
        onlyOwner 
        whenNotPaused 
    returns(
        uint256, // CDP total collateral
        uint256  // CDP total debt
    ) {
        CDP storage cdp = cdps[user];

        cdp.owner = user;
        cdp.collateral = add(cdp.collateral, collateral);

        int256 availableDebt = _calculateDebt(cdp);
        require(availableDebt >= 0, "Core/Negative available debt");
        require(uint256(availableDebt) >= debt, "Core/Exceed debt limit");
    
        cdp.debt = add(cdp.debt, debt);
        require(cdp.debt <= maxDebtCDP, "Core/Exceed CDP debt limit");
    
        _totalIncrease(collateral, debt);

        emit CDPFunded(user, collateral, debt);

        return (cdp.collateral, cdp.debt);
    }

    function defund(address user, uint256 collateral, uint256 debt) 
        external 
        onlyOwner 
        whenNotPaused 
    returns(
        uint256, // CDP 
        uint256  //
    ){
        CDP storage cdp = cdps[user];
        require(cdp.owner != address(0), "Core/CDP doesn't exist");

        require(cdp.debt > debt, "Core/Negative CDP debt");
        cdp.debt = sub(cdp.debt, debt);

        require(cdp.collateral > collateral, "Core/Negative CDP collateral");
        cdp.collateral = sub(cdp.collateral, collateral);

        int256 availableDebt = _calculateDebt(cdp);
        require(availableDebt >= 0, "Core/Negative CDP debt");
        
        _totalDecrease(collateral, debt);
        if (cdp.debt == 0 && cdp.collateral == 0) {
            emit CDPDeleted(user);
            delete cdps[user];
        }

        emit CDPDefunded(user, collateral, debt);

        return (cdp.collateral, cdp.debt);
    }

    function kickvidate(address liquidator, address user, uint256 liqdebt) 
        external 
        onlyOwner 
        whenNotPaused 
    returns(
        uint256 liquidatorCashback,
        uint256 liquidatorMargin,
        uint256 debtorCashback
    ) {
        CDP storage cdp = cdps[user];

        int256 debt = _calculateDebt(cdp);
        require(debt < 0, "Core/None zero debt");

        require(debt + int256(liqdebt) >= int256(cdp.debt), "Core/Negative liquidator cashback");
        liquidatorCashback = uint256(debt + int256(liqdebt) - int256(cdp.debt));

        uint256 totalCashback = _calculateCollateral(cdp);
        uint256 liquidatorFee = _calculateFee(totalCashback);

        debtorCashback = sub(totalCashback, liquidatorFee);
        liquidatorMargin = sub(cdp.collateral, debtorCashback);

        _totalDecrease(cdp.collateral, cdp.debt);
        delete cdps[user];

        emit CDPDeleted(user);
        emit CDPLiquidated(liquidator, user, liqdebt, liquidatorCashback, liquidatorMargin, debtorCashback);

        return (liquidatorCashback, liquidatorMargin, debtorCashback);
    }

    function _calculateFee(uint256 collateral) private view returns (uint256) {
        return wmul(collateral, liquidateFee);
    }

    function _calculateCollateral(CDP memory cdp) private view returns (uint256) {
        return sub(cdp.collateral, cdp.debt) / collateralPrice;
    }

    function _calculateDebt(CDP memory cdp) private view returns (int256) {
        return int256(wdiv(wmul(cdp.collateral, collateralPrice), collateralRate)) - int256(cdp.debt);
    }

    function _calculateCashback(uint256 l, uint256 r) private pure returns (int256 result) {
        result = int256(l) - int256(r);
        if (result >= 0) {
            return 0;
        }
        return -result;
    }

    function _totalIncrease(uint256 collateral, uint256 debt) private {
        totalCollateral = add(totalCollateral, collateral);
        totalDebt = add(totalDebt, debt);
        require(totalDebt < maxDebt, "Core/Debt reached limit");
    }

    function _totalDecrease(uint256 collateral, uint256 debt) private {
        totalCollateral = sub(totalCollateral, collateral);
        totalDebt = sub(totalDebt, debt);
    }
}