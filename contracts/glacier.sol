// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.18;
import "hardhat/console.sol";

import {IGov} from "./IGov.sol";
import {IZK} from "./IZK.sol";
import {parseProof} from "./zk_parser.sol";

contract Glacier {
    string public constant GLACIER_VERSION = "0.0.1";
    string public GLACIER_NETWORK;
    uint32 public latestBlock;
    address public gov;
    address public zk;

    constructor(
        string memory _network,
        address _gov,
        address _zk
    ) {
        gov = _gov;
        zk = _zk;
        GLACIER_NETWORK = _network;
    }

    struct BlockInfo {
        uint32 blockNumber;
        bytes32 blockHash;
        bytes32 preblockHash;
        uint256 timestamp;
        uint256[8] commitment; // block commitment
        string daid; // ar://<arid>?sha256=<hash>
    }

    mapping(uint32 => BlockInfo) public blockHub;

    event BlockCommit(uint32 blockNumber);

    function zkverifyOneBlock(BlockInfo memory _newBlock)
        internal
        returns (bool)
    {
        uint256[8] memory proof0 = _newBlock.commitment;
        uint256[2] memory a;
        uint256[2][2] memory b;
        uint256[2] memory c;

        (a, b, c) = parseProof(proof0);

        uint256[1] memory input = [
            uint256(_newBlock.blockHash)
        ];

        console.log("a0 %s, a1 %s", a[0], a[1]);
        console.log("b00 %s, b01 %s", b[0][0], b[0][1]);
        console.log("b10 %s, b11 %s", b[1][0], b[1][1]);
        console.log("c0 %s, c1 %s", c[0], c[1]);
        console.log("input %s", input[0]);

        return IZK(zk).verifyProof(a, b, c, input);
    }

    /// @notice Commit block
    function commitBlock(BlockInfo calldata _newBlockData) external {
        IGov(gov).requireActiveValidator(msg.sender);

        require(
            blockHub[_newBlockData.blockNumber].blockNumber !=
                _newBlockData.blockNumber,
            "block already exists"
        );

        // check latest block
        uint32 preblockNumber = _newBlockData.blockNumber - 1;
        if (preblockNumber != 0) {
            BlockInfo memory _lastCommittedBlock = blockHub[preblockNumber];
            require(
                _lastCommittedBlock.blockHash == _newBlockData.preblockHash,
                "preblockHash not match"
            );
            require(
                _lastCommittedBlock.timestamp <= _newBlockData.timestamp,
                "invalid timestamp"
            );
        }

        bool result = zkverifyOneBlock(_newBlockData);
        require(result, "invalid block zkcommitment");

        // store
        blockHub[_newBlockData.blockNumber] = _newBlockData;
        latestBlock = _newBlockData.blockNumber;

        // event
        emit BlockCommit(_newBlockData.blockNumber);
    }
}
