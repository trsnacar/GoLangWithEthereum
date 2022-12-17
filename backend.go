package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
)

// contractInstance is a global variable that holds a reference to the contract instance
var contractInstance *SimpleContract

func main() {
	// Connect to the Ethereum network
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	// Set the contract address
	contractAddress := common.HexToAddress("0x1234567890123456789012345678901234567890")

	// Create a new instance of the contract
	contractInstance, err = NewSimpleContract(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// Set up router and routes
	r := mux.NewRouter()
	r.HandleFunc("/balance", handleGetBalance).Methods("GET")
	r.HandleFunc("/deposit", handlePostDeposit).Methods("POST")

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", r))
}

func handleGetBalance(w http.ResponseWriter, r *http.Request) {
	// Get the contract balance
	balance, err := contractInstance.Balance(&bind.CallOpts{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the balance to the response
	fmt.Fprint(w, balance)
}

func handlePostDeposit(w http.ResponseWriter, r *http.Request) {
	// Parse the deposit amount from the request body
	var depositAmount big.Int
	if err := json.NewDecoder(r.Body).Decode(&depositAmount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Make a deposit
	tx, err := contractInstance.Deposit(auth, &depositAmount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Wait for the transaction to be mined
	ctx := context.Background()
	receipt, err := client.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the receipt to the response
	json.NewEncoder(w).Encode(receipt)
}
