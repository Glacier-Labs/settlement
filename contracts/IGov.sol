// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.0;

interface IGov {
    function requireActiveValidator(        
        address validator
    )
    external;
}