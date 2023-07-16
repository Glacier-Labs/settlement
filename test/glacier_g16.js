const {
  loadFixture,
} = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const { expect } = require("chai");
const mockBlocks = require('./mock_blocks.json');

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
        console.log(ethers.decodeBase64(block.blockHash))
        await glacier.commitBlock({
          blockNumber: block.blockNumber,
          blockHash: ethers.decodeBase64(block.blockHash),
          preblockHash: ethers.decodeBase64(block.preblockHash),
          timestamp: block.timestamp,
          commitment: ethers.decodeBase64(block.preblockHash),
          daid: block.daid,
        })
      }
    });
  });
});
