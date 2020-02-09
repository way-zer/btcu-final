package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

const CompositeKey = "name~author~press"

type RightChaincode struct {
}

func (t *RightChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Chaincode Init")
	return shim.Success(nil)
}

func (t *RightChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Chaincode Invoke")
	function, args := stub.GetFunctionAndParameters()
	if function == "register" {
		return t.register(stub, args)
	} else if function == "delete" {
		// Deletes an entity from its state
		// return t.delete(stub, args)
	} else if function == "query" {
		return t.query(stub, args)
	} else if function == "queryHash" {
		return t.queryHash(stub, args)
	}

	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}

// Transaction makes payment of X units from A to B
func (t *RightChaincode) register(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var Sha, Info string // Entities
	var err error

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	// Initialize the chaincode
	Sha = args[0]
	InfoKey, err := stub.CreateCompositeKey(CompositeKey, args[1:4])
	if err != nil {
		return shim.Error(err.Error())
	}
	Info = args[4]

	old, err := stub.GetState(Sha)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if old != nil {
		return shim.Error(fmt.Sprintf("Sha(%s) has registered with \"%s\"", Sha, old))
	}

	old2, err := stub.GetState(Sha)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if old2 != nil {
		return shim.Error(fmt.Sprintf("Info(%s) has registered with \"%s\"", InfoKey, old2))
	}

	// Write the state to the ledger
	err = stub.PutState(Sha, []byte(Info))
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf("Register %s with %s\n", Sha, Info)
	err = stub.PutState(InfoKey, []byte(Sha))
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf("Register %s with %s\n", InfoKey, Info)

	return shim.Success(nil)
}

// Deletes an entity from state
func (t *RightChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	A := args[0]

	// Delete the key from the state in ledger
	err := stub.DelState(A)
	if err != nil {
		return shim.Error("Failed to delete state")
	}

	return shim.Success(nil)
}

// query callback representing the query of a chaincode
func (t *RightChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	Sha := args[0]

	// Get the state from the ledger
	Info, err := stub.GetState(Sha)
	if err != nil {
		return shim.Error("Failed to get state for " + Sha)
	}
	fmt.Printf("Query %s Response:%s\n", Sha, Info)
	return shim.Success(Info)
}

func (t *RightChaincode) queryHash(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	InfoKey, err := stub.CreateCompositeKey(CompositeKey, args)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Get the state from the ledger
	Sha, err := stub.GetState(InfoKey)
	if err != nil {
		return shim.Error("Failed to get state for " + InfoKey)
	}
	fmt.Printf("Query %s Response:%s\n", InfoKey, Sha)
	return shim.Success(Sha)
}

func main() {
	err := shim.Start(new(RightChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
