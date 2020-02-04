package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Chaincode Init")
	_, args := stub.GetFunctionAndParameters()
	t.register(stub,args)

	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Chaincode Invoke")
	function, args := stub.GetFunctionAndParameters()
	if function == "register" {
		// Make payment of X units from A to B
		return t.register(stub, args)
	} else if function == "delete" {
		// Deletes an entity from its state
		// return t.delete(stub, args)
	} else if function == "query" {
		// the old "Query" is now implemtned in invoke
		return t.query(stub, args)
	}

	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}

// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) register(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var Sha,Info string    // Entities
	var err error

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	// Initialize the chaincode
	Sha = args[0]
	Info = args[1]

	old, err :=stub.GetState(Sha)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if old != nil {
		return shim.Error(fmt.Sprintf("Sha(%s) has registered with \"%s\"",Sha,old))
	}

	// Write the state to the ledger
	err = stub.PutState(Sha, []byte(Info))
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf("Register %s with %s\n", Sha, Info)

	return shim.Success(nil)
}

// Deletes an entity from state
func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
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
func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	Sha := args[0]

	// Get the state from the ledger
	Info, err := stub.GetState(Sha)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + Sha + "\"}"
		return shim.Error(jsonResp)
	}

	if Info == nil {
		jsonResp := "{\"Error\":\"Nil for " + Sha + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp := "{\"Sha\":\"" + Sha + "\",\"DATA\":\"" + string(Info) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return shim.Success([]byte(jsonResp))
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
