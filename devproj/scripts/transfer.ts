import { ethers } from "hardhat";

const CONTROLLER_ADDRESS = "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0";
const TOKEN_PRECISION = Math.pow(10, 18);

async function main() {
  const controller = await ethers.getContractAt("Controller", CONTROLLER_ADDRESS);
  const signers = await ethers.getSigners();

  await controller
    .connect(signers[0])
    .fund((5 * TOKEN_PRECISION).toLocaleString('fullwide', {useGrouping:false}));
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
