// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.18;

interface IZK {
    function verifyProof(        
        uint256[2] memory a,
        uint256[2][2] memory b,
        uint256[2] memory c,
        uint256[1] calldata input)
    external returns (bool);
}