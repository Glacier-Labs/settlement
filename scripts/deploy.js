// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// You can also run a script with `npx hardhat run <script>`. If you do that, Hardhat
// will compile your contracts, add the Hardhat Runtime Environment's members to the
// global scope, and execute the script.
const hre = require("hardhat");

async function main() {
  const network = 'glc001-testnet'
  const govAddr = '0xDDac83FA3C72276aAaf0045551c4F9FbAA35375a';
  const zkAddr = '0x80c748Cf77E9045F601a73df8A442045c7B4C4c7';
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
