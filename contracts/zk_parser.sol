// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.18;

function parseProof(bytes32 proof)
    pure
    returns (
        uint256[2] memory a,
        uint256[2][2] memory b,
        uint256[2] memory c
    )
{
    uint256[] memory els = splitBytes32(proof, 8);

    a = [els[0], els[1]];
    b = [[els[2], els[3]], [els[4], els[5]]];
    c = [els[6], els[7]];
    return (a, b, c);
}

function splitBytes32(bytes32 data, uint256 numElements)
    pure
    returns (uint256[] memory)
{
    uint256[] memory result = new uint256[](numElements);

    for (uint256 i = 0; i < numElements; i++) {
        uint256 element = uint256(data);
        result[i] = element;
        data = data >> 256; // Shift right by 256 bits
    }

    return result;
}
