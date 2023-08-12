// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// You can also run a script with `npx hardhat run <script>`. If you do that, Hardhat
// will compile your contracts, add the Hardhat Runtime Environment's members to the
// global scope, and execute the script.
const hre = require("hardhat");

async function main() {
  const network = 'glc001-testnet'
  const govAddr = '0xf10b40cE6C9071fA89649D4E1A87f07660eeF79f';
  const zkAddr = '0x9bF378F23C51b9faAf3541B1462A056BB0a7A087';
  const Glacier = await ethers.getContractFactory("Glacier");
  const glacier = await Glacier.deploy(network, govAddr, zkAddr);
  const glaciervAddr = await glacier.getAddress()

  console.log(
    `Glacier deployed to ${glaciervAddr}`
  );
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
