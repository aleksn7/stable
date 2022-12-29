import { ethers } from "hardhat";

const CORE_ADDRESS = "0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512";

async function main() {
  const core = await ethers.getContractAt("Core", CORE_ADDRESS);
  const signers = await ethers.getSigners();

  const price = await core
    .connect(signers[0])
    .collateralPrice();

  console.log(price);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
