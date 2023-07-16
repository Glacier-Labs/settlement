// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.18;

interface IGov {
    function requireActiveValidator(        
        address validator
    )
    external;
}