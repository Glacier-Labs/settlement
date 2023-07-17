const {
  loadFixture,
} = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const { loadProof, loadPublic } = require('./load_proof');

describe("GlacierVerifier", function () {
  async function deployFixture() {
    const [validator] = await ethers.getSigners();
    console.log(`validator ${validator.address}`)

    console.log(`deploy zk...`)
    const Zk = await ethers.getContractFactory("Verifier");
    const zk = await Zk.deploy();
    const zkAddr = await zk.getAddress()

    console.log(`deployed zk at ${zkAddr}`)
    return { zk };
  }

  describe("Deployment", function () {
    it("ZKProof", async function () {
      const { zk } = await loadFixture(deployFixture);

      const arr = loadProof('./test/proof.bin')

      let a = arr[0];
      let b = arr[1];
      let c = arr[2];
      let input = loadPublic('./test/public.bin')
      console.log(`a ${a}`)
      console.log(`b ${b}`)
      console.log(`c ${c}`)
      console.log(`input ${input}`)

      await zk.verifyProof(a, b, c, input)
    });
  });
});
