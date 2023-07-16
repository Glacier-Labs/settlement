require('dotenv').config();
require("@nomicfoundation/hardhat-toolbox");


OPBNB_PRIVATE_KEY = process.env.OPBNB_PRIVATE_KEY
TESTER_PRIVATE_KEY = process.env.TESTER_PRIVATE_KEY

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.18",
  networks: {
    opbnbtestnet: {
      url: `https://opbnb-testnet-rpc.bnbchain.org`,
      accounts: [OPBNB_PRIVATE_KEY]
    },
    nautchain: {
      url: `https://api.proteus.nautchain.xyz/solana`,
      accounts: [OPBNB_PRIVATE_KEY]
    },
    localhost: {
      accounts: [TESTER_PRIVATE_KEY]
    }
  },
};
