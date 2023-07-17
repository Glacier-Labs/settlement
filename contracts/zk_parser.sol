// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.18;

function parseProof(uint256[8] memory proof)
    pure
    returns (
        uint256[2] memory a,
        uint256[2][2] memory b,
        uint256[2] memory c
    )
{
    uint256[8] memory els = proof;

    a = [els[0], els[1]];
    b = [[els[2], els[3]], [els[4], els[5]]];
    c = [els[6], els[7]];

    return (a, b, c);
}
