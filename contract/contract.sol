pragma solidity ^0.7.0;

contract SimpleContract {
    uint public balance;

    constructor() public {
        balance = 100;
    }

    function deposit() public payable {
        require(msg.value > 0, "Deposit amount must be greater than 0");
        balance += msg.value;
    }

    function withdraw(uint amount) public {
        require(amount <= balance, "Insufficient balance");
        balance -= amount;
        msg.sender.transfer(amount);
    }
}
