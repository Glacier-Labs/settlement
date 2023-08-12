// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// You can also run a script with `npx hardhat run <script>`. If you do that, Hardhat
// will compile your contracts, add the Hardhat Runtime Environment's members to the
// global scope, and execute the script.
const hre = require("hardhat");
const { expect } = require("chai");
const { loadProof, loadPublic } = require('../test/load_proof');

async function main() {
  const zkAddr = '0xDDac83FA3C72276aAaf0045551c4F9FbAA35375a';

  const zk = await hre.ethers.getContractAt("Verifier", zkAddr);

  const arr = loadProof('./test/proof.bin')

  let a = arr[0];
  let b = arr[1];
  let c = arr[2];
  let input = loadPublic('./test/public.bin')
  console.log(`a ${a}`)
  console.log(`b ${b}`)
  console.log(`c ${c}`)
  console.log(`input ${input}`)

  console.log(await zk.verifyProof(a, b, c, input))
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});