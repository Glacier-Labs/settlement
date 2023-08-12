// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// You can also run a script with `npx hardhat run <script>`. If you do that, Hardhat
// will compile your contracts, add the Hardhat Runtime Environment's members to the
// global scope, and execute the script.
const hre = require("hardhat");
const { expect } = require("chai");
const mockBlocks = require('../test/mock_blocks.json');
const { parseProof } = require('../test/load_proof');

async function main() {
  const glacierAddr = '0xb16801c660A0777c3A2fE81a577De38071d6364F';

  const glacier = await hre.ethers.getContractAt("Glacier", glacierAddr);

  console.log(await glacier.GLACIER_VERSION())
  console.log(await glacier.GLACIER_NETWORK())
  expect(await glacier.GLACIER_VERSION()).to.equal('0.0.1');
  expect(await glacier.GLACIER_NETWORK()).to.equal('glc001-testnet');

  for (const block of mockBlocks) {
    // console.log(ethers.decodeBase64(block.blockHash))
    const arr = parseProof(block.commitment)
    const d = arr[3];
    await glacier.commitBlock({
      blockNumber: block.blockNumber,
      blockHash: block.blockHash,
      preblockHash: block.preblockHash,
      timestamp: block.timestamp,
      commitment: d,
      daid: block.daid,
    })
    break;
  }
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});