// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// You can also run a script with `npx hardhat run <script>`. If you do that, Hardhat
// will compile your contracts, add the Hardhat Runtime Environment's members to the
// global scope, and execute the script.
const hre = require("hardhat");

async function main() {
  const network = 'glc001-testnet'
  const govAddr = '0x3F878B8678BC6aFeDf3d3e6467DFfc38Fa7EFa97';
  const zkAddr = '0x7C6Cc25f0af38F8EC7Da9C2ff75c6e049891b2ac';
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
