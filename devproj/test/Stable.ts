import { time, loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { anyValue } from "@nomicfoundation/hardhat-chai-matchers/withArgs";
import { expect } from "chai";
import { ethers } from "hardhat";

const TOKEN_PRECISION = Math.pow(10, 18);
const RATE_PRECISION = Math.pow(10, 27);

async function deployCoin() {
  const Token = await ethers.getContractFactory("Coin");
  const token = await Token.deploy("Stable", "STB");
  await token.deployed();
  return token;
}

async function deployCore() {
  const MAX_DEBT = (50000 * TOKEN_PRECISION).toLocaleString('fullwide', {useGrouping:false});
  const MAX_DEBT_PER_CDP = (5000 * TOKEN_PRECISION).toLocaleString('fullwide', {useGrouping:false});
  const COLLATERAL_RATE = (2 * RATE_PRECISION).toLocaleString('fullwide', {useGrouping:false});
  const COLLATERAL_PRICE = (5 * RATE_PRECISION).toLocaleString('fullwide', {useGrouping:false});
  const LIQUIDATE_FEE = (2 * RATE_PRECISION / 10).toLocaleString('fullwide', {useGrouping:false});

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

    return {controller, core};
  }

  describe("Fund", function () {
    it("Should be right balance", async function () {
      const tests = [
        {
          collateral: BigInt(50 * TOKEN_PRECISION),
          debt: 0,
        },
        {
          collateral: BigInt(50 * TOKEN_PRECISION),
          debt: BigInt(50 * TOKEN_PRECISION),
        }
      ];
      
      const signers = await ethers.getSigners();
      const signer = signers[0];
      for (let test of tests) {
        const {controller, core} = await loadFixture(deployControlerFixture);

        await controller
          .connect(signer)
          .fund(test.debt, {value: test.collateral});

        const balance = await core.getBalance(signer.getAddress());
        expect(test.collateral).to.equal(balance[0]);
        expect(test.debt).to.equal(balance[1]);
      }
    });

    it("Should be reverted while funding", async function () {
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

    // it("Should set the right owner", async function () {
    //   const { lock, owner } = await loadFixture(deployOneYearLockFixture);

    //   expect(await lock.owner()).to.equal(owner.address);
    // });

    // it("Should receive and store the funds to lock", async function () {
    //   const { lock, lockedAmount } = await loadFixture(
    //     deployOneYearLockFixture
    //   );

    //   expect(await ethers.provider.getBalance(lock.address)).to.equal(
    //     lockedAmount
    //   );
    // });

    // it("Should fail if the unlockTime is not in the future", async function () {
    //   // We don't use the fixture here because we want a different deployment
    //   const latestTime = await time.latest();
    //   const Lock = await ethers.getContractFactory("Lock");
    //   await expect(Lock.deploy(latestTime, { value: 1 })).to.be.revertedWith(
    //     "Unlock time should be in the future"
    //   );
    // });
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
        const {controller, core} = await loadFixture(deployControlerFixture);
        
        await controller
          .connect(signer)
          .fund(test.debt, {value: test.collateral});

        const balanceBefore = await signer.getBalance();
        await controller
          .connect(signer)
          .defund(test.defundCollateral, test.defundDebt);
        const balanceAfter = await signer.getBalance();

        const balance = await core.getBalance(signer.getAddress());
        expect(BigInt(TOKEN_PRECISION)).to.greaterThan(
          test.defundCollateral - (balanceAfter.toBigInt() - balanceBefore.toBigInt())
        );
        expect(test.collateral - test.defundCollateral).to.equal(balance[0]);
        expect(test.debt -  test.defundDebt).to.equal(balance[1]);
      }
    });

    it("Should be reverted while defunding", async function () {
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

      // it("Should revert with the right error if called from another account", async function () {
      //   const { lock, unlockTime, otherAccount } = await loadFixture(
      //     deployOneYearLockFixture
      //   );

      //   // We can increase the time in Hardhat Network
      //   await time.increaseTo(unlockTime);

      //   // We use lock.connect() to send a transaction from another account
      //   await expect(lock.connect(otherAccount).withdraw()).to.be.revertedWith(
      //     "You aren't the owner"
      //   );
      // });

      // it("Shouldn't fail if the unlockTime has arrived and the owner calls it", async function () {
      //   const { lock, unlockTime } = await loadFixture(
      //     deployOneYearLockFixture
      //   );

      //   // Transactions are sent using the first signer by default
      //   await time.increaseTo(unlockTime);

      //   await expect(lock.withdraw()).not.to.be.reverted;
      // });

  //   describe("Events", function () {
  //     it("Should emit an event on withdrawals", async function () {
  //       const { lock, unlockTime, lockedAmount } = await loadFixture(
  //         deployOneYearLockFixture
  //       );

  //       await time.increaseTo(unlockTime);

  //       await expect(lock.withdraw())
  //         .to.emit(lock, "Withdrawal")
  //         .withArgs(lockedAmount, anyValue); // We accept any value as `when` arg
  //     });
  //   });

  //   describe("Transfers", function () {
  //     it("Should transfer the funds to the owner", async function () {
  //       const { lock, unlockTime, lockedAmount, owner } = await loadFixture(
  //         deployOneYearLockFixture
  //       );

  //       await time.increaseTo(unlockTime);

  //       await expect(lock.withdraw()).to.changeEtherBalances(
  //         [owner, lock],
  //         [lockedAmount, -lockedAmount]
  //       );
  //     });
  //   });
  });
});
