// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.18;


contract GlacierGov {
    address[] validators;
    address public owner;

    constructor(address root) {
        validators.push(root);
        owner = msg.sender;
    }

    function requireActiveValidator(        
        address validator
    ) view external {
        require(validators[0] == validator, "403");
    }

    // TODO https://github.com/OpenZeppelin/openzeppelin-contracts/tree/master/contracts/governance
}