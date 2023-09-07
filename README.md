# settlement

A Contract to prove the block's DA. Any validator can submit a valid block. The contract will verify the block's 

    1. validator's sign
    2. zk proof for txs's final state
    3. commitment
    4. governance selects the validator

Everyone can verify it's valid. If invalid, slash the validator.

Just Verify without Execute, powered by zk.

## Data Struct

block:

    - blockNumber
    - blockHash
    - preblockHash
    - commitment // zk hash proof
    - timestamp
    - sign
    - daid

daid protocol:

    - ar://<arid>?sha256=<hash_of_data>
    - gfn://<bucket/object_key>?sha256=<hash_of_data>
    - ipfs://<cid>?sha256=<hash_of_data>

## ZK Proof

- blockpreimage = sha256(blockData)
- blockHash = mimcHash(blocknumber, blockpreimage, preblockhash, timestamp)
- sign = zksign(validator, blockHash)
- commitment = zkProof(blocknumber, blockpreimage, preblockhash, timestamp, sign)

    - public: 
        - blochHash

## Verify

- require(zkproof(commitment, blockhash))
- require(sha256(BlockData(DaId)) == "BlockPreImage")



## Contract

- Pre-requisites

```
npm i
```

- Configure & Deploy

```
npx hardhat node
export TESTER_PRIVATE_KEY=<PrivateKey>
npx hardhat run scripts/deploy_zk.js --network localhost
npx hardhat run scripts/deploy_gov.js --network localhost
npx hardhat run scripts/deploy.js --network localhost
```

pls check the `scripts/*` for more details

- Test

```
npx hardhat test test/glacier_g16.js --network localhost
npx hardhat test test/zk_g16.js --network localhost
```

- opbnb testnet
    - ZK: 0x7C6Cc25f0af38F8EC7Da9C2ff75c6e049891b2ac
    - Gov: 0x3F878B8678BC6aFeDf3d3e6467DFfc38Fa7EFa97
    - GlacierV1: 0x4c63fa561D39a7E91Dd75cA75c29a5a0607157c8

- verify contract on opbnb testnet

```
npx hardhat run --network opbnbtestnet scripts/verify_contract.js
Successfully submitted source code for contract
contracts/glacier.sol:Glacier at 0x4c63fa561D39a7E91Dd75cA75c29a5a0607157c8
for verification on the block explorer. Waiting for verification result...

Successfully verified contract Glacier on the block explorer.
https://opbnbscan.com/address/0x4c63fa561D39a7E91Dd75cA75c29a5a0607157c8#code

npx hardhat run --network opbnbtestnet scripts/verify_contract_zk.js
Successfully submitted source code for contract
contracts/zkglacier_g16.sol:Verifier at 0x7C6Cc25f0af38F8EC7Da9C2ff75c6e049891b2ac
for verification on the block explorer. Waiting for verification result...

Successfully verified contract Verifier on the block explorer.
https://opbnbscan.com/address/0x7C6Cc25f0af38F8EC7Da9C2ff75c6e049891b2ac#code
```