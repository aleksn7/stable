import { ethers } from "hardhat";
import { Contract } from "@ethersproject/contracts/src.ts";

const TOKEN_PRECISION = Math.pow(10, 18);
const RATE_PRECISION = Math.pow(10, 27);

async function deployCoin() {
  const Token = await ethers.getContractFactory("Coin");
  const token = await Token.deploy("Stable", "STB");
  await token.deployed();
  return token;
}

async function deployCore() {
  const MAX_DEBT = BigInt(50000 * TOKEN_PRECISION);
  const MAX_DEBT_PER_CDP = BigInt(5000 * TOKEN_PRECISION);
  const COLLATERAL_RATE = BigInt(2 * (RATE_PRECISION));
  const COLLATERAL_PRICE = BigInt(5 * RATE_PRECISION / 10);
  const LIQUIDATE_FEE = BigInt(2 * RATE_PRECISION / 10);

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

async function main() {
  const coin = await deployCoin();
  console.log(`Coin deployed at address: ${coin.address}`);

  const core = await deployCore();
  console.log(`Core deployed at address: ${core.address}`);

  const controller = await deployController(core.address, coin.address);
  console.log(`Controller deployed at address: ${controller.address}`);

  await coin.transferOwnership(controller.address);
  await core.transferOwnership(controller.address);

  console.log(`Coin and core ownership transfered to controller`);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
