// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// You can also run a script with `npx hardhat run <script>`. If you do that, Hardhat
// will compile your contracts, add the Hardhat Runtime Environment's members to the
// global scope, and execute the script.
const hre = require("hardhat");
const { expect } = require("chai");

async function main() {
  const zkContractAddr = '0xDDac83FA3C72276aAaf0045551c4F9FbAA35375a';

  const verifier = await hre.ethers.getContractAt("GlacierVerifier", zkContractAddr);

  console.log(await verifier.GLACIER_VERSION())
  console.log(await verifier.GLACIER_NETWORK())
  expect(await verifier.GLACIER_VERSION()).to.equal('0.0.1');
  expect(await verifier.GLACIER_NETWORK()).to.equal('glc001-testnet');

  let a = [BigInt(`11370649296748688127075619763598486259227544270632091037082135364618695166743`), BigInt(`979592821429824271516747744716510753562925713858944572553381163419647959381`)];
  let b = [[BigInt(`21264786916418983771692915523337765456466536515884782365057660304912218013657`), BigInt(`15615050179246187606291120331568975309267663327306326569152216801249987963373`)], [BigInt(`15867751724428410037851902131601060859896682890887527857132011706113466915144`), BigInt(`18419037024947252377409852105519604948380245269684616868754275791688114403515`)]];
  let c = [BigInt(`11347903346549542389633446095643356647789279909104580929688581024560638003923`), BigInt(`17181928967565418095669845539765054019487703098671937222050284275716385344280`)];
  let input = [BigInt(`7850748157841143743037130756160722548816244036246230200666437080116841652363`)];

  let result = await verifier.verifyProof(a, b, c, input)
  console.log(`zk proof: ${result}`)
  expect(result).to.equal(true)

  let invaildInput = [BigInt(`1850748157841143743037130756160722548816244036246230200666437080116841652363`)];
  result = await verifier.verifyProof(a, b, c, invaildInput)

  console.log(`zk proof: ${result}`)
  expect(result).to.equal(false)
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});