import { time, loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { anyValue } from "@nomicfoundation/hardhat-chai-matchers/withArgs";
import { expect } from "chai";
import { ethers } from "hardhat";
import { any } from "hardhat/internal/core/params/argumentTypes";

const TOKEN_PRECISION = Math.pow(10, 18);
const RATE_PRECISION = Math.pow(10, 18);

const MAX_DEBT = BigInt(50000 * TOKEN_PRECISION);
const MAX_DEBT_PER_CDP = BigInt(5000 * TOKEN_PRECISION);
const COLLATERAL_RATE = BigInt(2 * RATE_PRECISION);
const COLLATERAL_PRICE = BigInt(5 * RATE_PRECISION);
const LIQUIDATE_FEE = BigInt(2 * RATE_PRECISION / 10);

async function deployCoin() {
  const Token = await ethers.getContractFactory("Coin");
  const token = await Token.deploy("Stable", "STB");
  await token.deployed();
  return token;
}

async function deployCore() {
  const Core = await ethers.getContractFactory("Core");
  const core = await Core.deploy(
    MAX_DEBT, MAX_DEBT_PER_CDP, COLLATERAL_RATE, COLLATERAL_PRICE, LIQUIDATE_FEE);

  await core.deployed();
  return core;
}

async function deployController(coreAddress: string, coinAddress: string) {
  const Controller = await ethers.getContractFactory("Controller");
  const controller = await Controller.deploy(coreAddress, coinAddress);
  await controller.deployed();
  return controller;
}

describe("Stable", function () {
  // We define a fixture to reuse the same setup in every test.
  // We use loadFixture to run this setup once, snapshot that state,
  // and reset Hardhat Network to that snapshot in every test.
  async function deployControlerFixture() {
    const coin = await deployCoin();
    const core = await deployCore();
    const controller = await deployController(core.address, coin.address);
  
    await coin.transferOwnership(controller.address);
    await core.transferOwnership(controller.address);

    return {controller, core, coin};
  }

  describe("Fund", function () {
    it("Should be funded", async function () {
      const tests = [
        {
          collateral: BigInt(50 * TOKEN_PRECISION),
          debt: 0,
        },
        {
          collateral: BigInt(50 * TOKEN_PRECISION),
          debt: BigInt(125 * TOKEN_PRECISION),
        }
      ];
      
      const signers = await ethers.getSigners();
      const signer = signers[0];
      for (let test of tests) {
        const {controller, core, coin} = await loadFixture(deployControlerFixture);
        
        await expect(
          controller
            .connect(signer)
            .fund(test.debt, {value: test.collateral})
        ).to.changeTokenBalance(coin, signer, test.debt)
         .to.changeEtherBalance(signer, -test.collateral)
         .to.emit(core, "CDPFunded").withArgs(signer.address, test.collateral, test.debt);
        expect(
          (await core.getBalance(signer.getAddress())).map((elem) => {
            return elem.toBigInt();
          })
        ).to.deep.equal([test.collateral, test.debt]);
      }
    });

    it("Should be reverted", async function () {
      const tests = [
        {
          collateral: BigInt(0),
          debt: BigInt(50 * TOKEN_PRECISION),
          reason: 'Core/Exceed debt limit'
        },
        {
          collateral: BigInt(50 * TOKEN_PRECISION),
          debt: BigInt(126 * TOKEN_PRECISION),
          reason: 'Core/Exceed debt limit'
        },
        {
          collateral: BigInt(5000 * TOKEN_PRECISION),
          debt: BigInt(6000 * TOKEN_PRECISION),
          reason: "Core/Exceed CDP debt limit"
        }
      ];
      
      const signers = await ethers.getSigners();
      const signer = signers[0];
      for (let test of tests) {
        const {controller} = await loadFixture(deployControlerFixture);

        await expect(
          controller
            .connect(signer)
            .fund(test.debt, {value: test.collateral})
        ).to.be.revertedWith(test.reason);
      }
    });
  });

  describe("Defund", function () {
    it("Should be defunded", async function () {
      const tests = [
        {
          collateral: BigInt(50 * TOKEN_PRECISION),
          debt: BigInt(125 * TOKEN_PRECISION),
          defundCollateral: BigInt(25 * TOKEN_PRECISION),
          defundDebt: BigInt(100 * TOKEN_PRECISION),
        },
      ];
      
      const signers = await ethers.getSigners();
      const signer = signers[0];
      for (let test of tests) {
        const {controller, core, coin} = await loadFixture(deployControlerFixture);
        
        await controller
          .connect(signer)
          .fund(test.debt, {value: test.collateral});

        await expect(
          controller
            .connect(signer)
            .defund(test.defundCollateral, test.defundDebt)
        ).to.changeTokenBalance(coin, signer, -test.defundDebt)
         .to.changeEtherBalances([signer, controller], [test.defundCollateral, -test.defundCollateral])
         .to.emit(core, "CDPDefunded").withArgs(signer.address, test.defundCollateral, test.defundDebt);

        expect(
          (await core.getBalance(signer.getAddress())).map((elem) => {
            return elem.toBigInt();
          })
        ).to.deep.equal([test.collateral - test.defundCollateral, test.debt -  test.defundDebt]);
      }
    });

    it("Should be reverted", async function () {
      const tests = [
        {
          collateral: BigInt(50 * TOKEN_PRECISION),
          debt: BigInt(125 * TOKEN_PRECISION),
          defundCollateral: BigInt(100 * TOKEN_PRECISION),
          defundDebt: BigInt(100 * TOKEN_PRECISION),
          reason: "Core/Negative CDP collateral"
        },
        {
          collateral: BigInt(50 * TOKEN_PRECISION),
          debt: BigInt(125 * TOKEN_PRECISION),
          defundCollateral: BigInt(49 * TOKEN_PRECISION),
          defundDebt: BigInt(100 * TOKEN_PRECISION),
          reason: "Core/Negative CDP debt"
        },
      ];
      
      const signers = await ethers.getSigners();
      const signer = signers[0];
      for (let test of tests) {
        const {controller} = await loadFixture(deployControlerFixture);
        
        await controller
          .connect(signer)
          .fund(test.debt, {value: test.collateral});
        await expect(
          controller
            .connect(signer)
            .defund(test.defundCollateral, test.defundDebt)
        ).to.revertedWith(test.reason);
      }
    });
  });
  describe("Liquidate", function () {
    it("Should be liquidated", async function () {
      const tests = [
        {
          collateral: BigInt(50 * TOKEN_PRECISION),
          debt: BigInt(125 * TOKEN_PRECISION),
          newPrice: COLLATERAL_PRICE / BigInt(2),
          liqDebt: BigInt(625 * TOKEN_PRECISION / 10) + BigInt(125 * TOKEN_PRECISION),
        }
      ];
      
      const signers = await ethers.getSigners();
      const owner = signers[0];
      const debtor = signers[0];
      const liquidator = signers[1];
      for (let test of tests) {
        const {controller, core, coin} = await loadFixture(deployControlerFixture);

        await controller
          .connect(debtor)
          .fund(test.debt, {value: test.collateral});
        await controller
          .connect(liquidator)
          .fund(test.liqDebt, {value: BigInt(2) * test.collateral});
        await controller
          .connect(owner)
          .setCollateralPrice(test.newPrice);
        
        await expect(
          controller
            .connect(liquidator)
            .kick(debtor.address, test.liqDebt)
        ).to.changeTokenBalance(coin, liquidator, -test.liqDebt)
          .to.changeEtherBalance(liquidator, BigInt(50 * TOKEN_PRECISION))
          .to.emit(core, "CDPDeleted").withArgs(debtor.address)
          .to.emit(core, "CDPLiquidated").withArgs(liquidator.address, debtor.address, test.liqDebt, anyValue, anyValue, anyValue);
      }
    });

    it("Should be reverted", async function () {
      const tests = [
        {
          collateral: BigInt(50 * TOKEN_PRECISION),
          debt: BigInt(125 * TOKEN_PRECISION),
          newPrice: COLLATERAL_PRICE / BigInt(2),
          liqDebt: BigInt(625 * TOKEN_PRECISION / 10) + BigInt(125 * TOKEN_PRECISION),
          reason: "Core/None debt here"
        }
      ];
      
      const signers = await ethers.getSigners();
      const debtor = signers[0];
      const liquidator = signers[1];
      for (let test of tests) {
        const {controller} = await loadFixture(deployControlerFixture);
        await controller
          .connect(debtor)
          .fund(test.debt, {value: test.collateral});
        await controller
          .connect(liquidator)
          .fund(test.liqDebt, {value: BigInt(2) * test.collateral});
        await expect(
          controller
            .connect(liquidator)
            .kick(debtor.address, test.liqDebt)
        ).to.revertedWith(test.reason);
      }
    });
  });
});
