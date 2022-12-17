package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Connect to the Ethereum network
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	// Set the contract address
	contractAddress := common.HexToAddress("0x1234567890123456789012345678901234567890")

	// Create a new instance of the contract
	contract, err := NewSimpleContract(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// Get the contract balance
	balance, err := contract.Balance(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Contract balance:", balance)

	// Make a deposit
	depositAmount := big.NewInt(100000000000000000) // 0.1 Ether
	tx, err := contract.Deposit(auth, depositAmount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deposit transaction:", tx.Hash().Hex())

	// Wait for the transaction to be mined
	ctx := context.Background()
	receipt, err := client.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Transaction mined in block:", receipt.BlockNumber)

	// Check the contract balance again
	balance, err = contract.Balance(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Contract balance:", balance)
}
