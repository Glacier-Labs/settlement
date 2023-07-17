const {
  loadFixture,
} = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const { expect } = require("chai");
const mockBlocks = require('./mock_blocks.json');
const { parseProof } = require('./load_proof');

describe("GlacierVerifier", function () {
  async function deployFixture() {
    const [validator] = await ethers.getSigners();
    console.log(`validator ${validator.address}`)
    const network = 'glc001-testnet'


    const Gov = await ethers.getContractFactory("GlacierGov");
    console.log(`deploy gov...`)
    const gov = await Gov.deploy(validator.address);
    const govAddr = await gov.getAddress()
    console.log(`deployed gov at ${govAddr}`)

    console.log(`deploy zk...`)
    const Zk = await ethers.getContractFactory("Verifier");
    const zk = await Zk.deploy();
    const zkAddr = await zk.getAddress()

    console.log(`deployed zk at ${zkAddr}`)

    const Glacier = await ethers.getContractFactory("Glacier");
    const glacier = await Glacier.deploy(network, govAddr, zkAddr);

    return { glacier };
  }

  describe("Deployment", function () {
    it("ZKProof", async function () {
      const { glacier } = await loadFixture(deployFixture);
      console.log(await glacier.GLACIER_VERSION())
      console.log(await glacier.GLACIER_NETWORK())
      expect(await glacier.GLACIER_VERSION()).to.equal('0.0.1');
      expect(await glacier.GLACIER_NETWORK()).to.equal('glc001-testnet');


      for (const block of mockBlocks) {
        // console.log(ethers.decodeBase64(block.blockHash))

        const arr = parseProof(block.commitment)
        const d = arr[3];
        // console.log("d:", d)
        console.log(block.blockNumber, block.blockHash)
        await glacier.commitBlock({
          blockNumber: block.blockNumber,
          blockHash: block.blockHash,
          preblockHash: block.preblockHash,
          timestamp: block.timestamp,
          commitment: d,
          daid: block.daid,
        })
      }
    });
  });
});
