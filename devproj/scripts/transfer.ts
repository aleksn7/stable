import { ethers } from "hardhat";

const MARKET_ADDRESS = "0x1f10F3Ba7ACB61b2F50B9d6DdCf91a6f787C0E82";

async function main() {
  const market = await ethers.getContractAt("Market", MARKET_ADDRESS);
  const signers = await ethers.getSigners();

  market
    .connect(signers[1])
    .mint({ value: `100` });
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
