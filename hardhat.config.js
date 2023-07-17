require('dotenv').config();
require("@nomicfoundation/hardhat-toolbox");


OPBNB_PRIVATE_KEY = process.env.OPBNB_PRIVATE_KEY
TESTER_PRIVATE_KEY = process.env.TESTER_PRIVATE_KEY

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: {
    version: "0.8.18",
    "evmVersion": "london",
    settings: {
      "optimizer": {
        "enabled": false,
        "runs": 200
      },
      "debug": {
        // How to treat revert (and require) reason strings. Settings are
        // "default", "strip", "debug" and "verboseDebug".
        // "default" does not inject compiler-generated revert strings and keeps user-supplied ones.
        // "strip" removes all revert strings (if possible, i.e. if literals are used) keeping side-effects
        // "debug" injects strings for compiler-generated internal reverts, implemented for ABI encoders V1 and V2 for now.
        // "verboseDebug" even appends further information to user-supplied revert strings (not yet implemented)
        "revertStrings": "debug"
      },
      "evmVersion": "london",
      "outputSelection": {
        "*": {
          "*": [
            "abi",
            "evm.bytecode.object",
            "evm.bytecode.opcodes",
            "evm.deployedBytecode",
            "evm.methodIdentifiers",
            "metadata"
          ],
          "": [
            "ast"
          ]
        }
      }
    }
  },
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
