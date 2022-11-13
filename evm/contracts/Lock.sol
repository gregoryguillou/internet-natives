// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

// Uncomment this line to use console.log
// import "hardhat/console.sol";

contract Lock {
    uint public unlockTime;
    address payable public owner;

    event Withdrawal(uint amount, uint when);

    constructor(uint _unlockTime) payable {
        require(
            block.timestamp < _unlockTime,
            "Unlock time should be in the future"
        );

        unlockTime = _unlockTime;
        owner = payable(msg.sender);
    }

    // Create a modifier that the function can be execulted only after unlockTime
    modifier isUnlocked() {
        require(block.timestamp >= unlockTime, "You can't withdraw yet");
        _;
    }

    // Create a modifier that the msg.sender must be the owner modifier
    modifier onlyOwner() {
        require(msg.sender == owner, "You aren't the owner");
        _;
    }

    function withdraw() public onlyOwner isUnlocked {
        // Uncomment this line, and the import of "hardhat/console.sol", to print a log in your terminal
        // console.log("Unlock time is %o and block timestamp is %o", unlockTime, block.timestamp);
        emit Withdrawal(address(this).balance, block.timestamp);
        owner.transfer(address(this).balance);
    }
}
