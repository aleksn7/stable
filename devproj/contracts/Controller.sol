// SPDX-License-Identifier: WTFPL
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/access/Ownable.sol";
// import "hardhat/console.sol";

interface CoreLike {
    function setCollateralPrice(uint256 price) external;

    function fund(address user, uint256 collateral, uint256 debt) external returns(
        uint256, // collateral
        uint256  // debt
    );

    function defund(address user, uint256 collateral, uint256 debt) external returns(
        uint256,  // collateral
        uint256   // debt
    );

    function liquidate(address liquidator, address user, uint256 liqdebt) external returns(
        uint256 liquidatorCashback,
        uint256 liquidatorMargin,
        uint256 debtorCashback
    );
}   

interface CoinLike {
    function mint(address account, uint256 amount) external;
    function burn(address account, uint256 amount) external;
    function balanceOf(address account) external view returns (uint256);
}

contract Controller is Ownable {
    CoreLike private _core;
    CoinLike private _coin;

    constructor(address core, address coin) {
        setCore(core);
        setCoin(coin);
    }

    function setCore(address core) public onlyOwner {
        _core = CoreLike(core);
    }

    function setCoin(address coin) public onlyOwner {
        _coin = CoinLike(coin);
    }

    function setCollateralPrice(uint256 price) external {
        _core.setCollateralPrice(price);
    }

    function fund(uint256 debt) external payable {
        address sender = msg.sender;
        uint256 collateral = msg.value;

        _core.fund(sender, collateral, debt);
        _coin.mint(sender, debt);
    }

    function defund(uint256 collateral, uint256 debt) external payable {
        address sender = msg.sender;
        uint256 balance = _coin.balanceOf(sender);
        require(balance >= debt, "Controller/Insufficient balance");

        _core.defund(sender, collateral, debt);
        _coin.burn(sender, debt);
        payable(sender).transfer(collateral);
    }

    function kick(address user, uint256 liqdebt) external payable {
        address sender = msg.sender;
        uint256 balance = _coin.balanceOf(sender);
        require(balance >= liqdebt, "Controller/Insufficient balance");

        uint256 liquidatorCashback;
        uint256 liquidatorMargin;
        uint256 debtorCashback;
        (liquidatorCashback, liquidatorMargin, debtorCashback) = _core.liquidate(sender, user, liqdebt);

        require(liqdebt >= debtorCashback, "Controller/Negative debtor");

        payable(sender).transfer(liquidatorMargin);
        _coin.burn(sender, liqdebt - debtorCashback);
        payable(user).transfer(debtorCashback);
    }
}


 