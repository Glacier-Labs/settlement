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

- opbnb
    - ZK: 0x7C6Cc25f0af38F8EC7Da9C2ff75c6e049891b2ac
    - Gov: 0x3F878B8678BC6aFeDf3d3e6467DFfc38Fa7EFa97
    - GlacierV1: 0x4c63fa561D39a7E91Dd75cA75c29a5a0607157c8