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

- Nauilus
    - ZK: 0x9bF378F23C51b9faAf3541B1462A056BB0a7A087
    - ZK 0.8.0: 0xa5f5C6D721A0DcCf2a0DDf03322EAa23bdbA653A
    - Gov: 0xf10b40cE6C9071fA89649D4E1A87f07660eeF79f
    - GLacierV1: 0x605caeeC07E114Ab741034b0C06385B1a588dBe9
    - GlacierV2: 0xb16801c660A0777c3A2fE81a577De38071d6364F
    - ZK 0.8.0: 0xa5f5C6D721A0DcCf2a0DDf03322EAa23bdbA653A
    - 0xd82b20EC429E9B671861e073a74047983EDfb22B
    - 0xec7DADc6d571fb26db2E2E11A31cec8F5E23EE1b
    - ZK 0.8.14: 0xbe6E0Cf9f127D70A4aE06345AA47b34e203EeD0e
    - ZK 0.8.10: 0xE6835CBD67073C47Ee10bc70DeF0Fa050516c090

- https://www.eclipse.builders/
    - https://docs.eclipse.builders/building-on-eclipse/documentation-for-eclipse-with-evm



curl -X POST https://faucet.evm.eclipse.eclipsenetwork.xyz/request_neon 'Content-Type: application/json' -d '{ "wallet": "0x830f93b43a737a7b45d84b1631c58e8fe54d0afc", "amount": 1 }'
https://api.evm.apricot.eclipsenetwork.xyz/solana
- ZK: 0xDDac83FA3C72276aAaf0045551c4F9FbAA35375a