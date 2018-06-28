/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

//WARNING - this chaincode's ID is hard-coded in chaincode_example04 to illustrate one way of
//calling chaincode from a chaincode. If this example is modified, chaincode_example04.go has
//to be modified as well with the new ID of chaincode_example02.
//chaincode_example05 show's how chaincode ID can be passed in as a parameter instead of
//hard-coding.

import (
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"time"
	"math/rand"
	"encoding/json"
)

var start time.Time
var countMess int

//custom data models
type AirforceComm struct {
	EntityID string `json:"entityID"`
	EntityVal string `json:"entityVal"`
	Message string `json:"message"`
}

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Init")
	_, args := stub.GetFunctionAndParameters()
	var A, B, mess, messagStr string    // Entities
	//var Aval, Bval int // Asset holdings
	var err error

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	// Initialize the chaincode
	A = args[0]
	B = args[1]

	mess = args[2]
	fmt.Printf("Aval = %d, Bval = %d\n, mess = %s/n", A, B, mess)

	// Write the state to the ledger
	messagStr = "{\"entityID\":" + A + ",\"entityVal\":" + B + ",\"message\":"+mess+"}"
	err = stub.PutState(A, []byte(messagStr))
	if err != nil {
		return shim.Error(err.Error())
	}

	messagStr = "{\"entityID\":" + B + ",\"entityVal\":" + A + ",\"message\":"+mess+"}"
	err = stub.PutState(B, []byte(messagStr))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Invoke")
	function, args := stub.GetFunctionAndParameters()
	if function == "invoke" {
		// Make payment of X units from A to B
		return t.invoke(stub, args)
	} else if function == "delete" {
		// Deletes an entity from its state
		return t.delete(stub, args)
	} else if function == "query" {
		// the old "Query" is now implemtned in invoke
		return t.query(stub, args)
	} else if function == "add" {
		return t.add(stub, args)
	} else if function == "rand" {
		return t.rand(stub, args)
	}

	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}

//Adds a new ID entity to the blockchain
func (t *SimpleChaincode) add(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var A, B, mess, messagStr string    // Entities
	//var Aval, Bval int // Asset holdings
	//var X int          // Transaction value
	var err error

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	A = args[0]
	B = args[1]
	mess = args[2]

	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger

	messagStr = "{\"entityID\":" + A + ",\"entityVal\":" + B + ",\"message\":"+mess+"}"

	err = stub.PutState(A, []byte(messagStr))
	if err != nil {
		return shim.Error("error with A PutState")
	}

	messagStr = "{\"entityID\":" + B + ",\"entityVal\":" + A + ",\"message\":"+mess+"}"
	err = stub.PutState(B, []byte(messagStr))
	if err != nil {
		return shim.Error("error with B PutState")
	}

	return shim.Success(nil)
}

// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) perform(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var A, B, mess string    // Entities
	//var Aval, Bval int // Asset holdings
	//var X int          // Transaction value
	var C int
	var err error

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	A = args[0]
	B = args[1]
	mess = args[2]
	C, err = strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("error strconverting arg2")
	}

	countMess = C
	mess = strconv.Itoa(countMess)

	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger

	Avalbytes, err := stub.GetState(A)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Avalbytes == nil {
		return shim.Error("Entity not found")
	}

	airStruc := AirforceComm{}
	json.Unmarshal(Avalbytes, &airStruc)
	airStruc.EntityID = A
	airStruc.EntityVal = B
	airStruc.Message = mess
	//Aval, _ = strconv.Atoi(string(Avalbytes))

	fmt.Printf("Aval = %s, Bval = %s, mess = %s\n", airStruc.EntityID, airStruc.EntityVal, airStruc.Message)

	start = time.Now()
	updMessage, _ := json.Marshal(airStruc)
	// Write the state back to the ledger
	err = stub.PutState(A, updMessage)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) rand(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var A, B, mess string    // Entities
	//var Aval, Bval int // Asset holdings
	//var X int          // Transaction value
	var C int
	var err error

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	A = args[0]
	B = args[1]
	mess = args[2]
	C, err = strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("error strconverting arg2")
	}

	mess = RandStringBytes(C)

	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger

	Avalbytes, err := stub.GetState(A)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Avalbytes == nil {
		return shim.Error("Entity not found")
	}

	airStruc := AirforceComm{}
	json.Unmarshal(Avalbytes, &airStruc)
	airStruc.EntityID = A
	airStruc.EntityVal = B
	airStruc.Message = mess
	//Aval, _ = strconv.Atoi(string(Avalbytes))

	fmt.Printf("Aval = %s, Bval = %s, mess = %s\n", airStruc.EntityID, airStruc.EntityVal, airStruc.Message)

	updMessage, _ := json.Marshal(airStruc)
	// Write the state back to the ledger
	err = stub.PutState(A, updMessage)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var A, B, mess string    // Entities
	//var Aval, Bval int // Asset holdings
	//var X int          // Transaction value
	var err error

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	A = args[0]
	B = args[1]
	mess = args[2]

	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger

	Avalbytes, err := stub.GetState(A)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Avalbytes == nil {
		fmt.Printf("run alert function for suspicious id", A)
		e_string := "run alert function for suspicious id" + A
		return shim.Error(e_string)
	}
	airStruc := AirforceComm{}
	json.Unmarshal(Avalbytes, &airStruc)
	airStruc.EntityID = A
	airStruc.EntityVal = B
	airStruc.Message = mess
	//Aval, _ = strconv.Atoi(string(Avalbytes))

	fmt.Printf("Aval = %s, Bval = %s, mess = %s\n", airStruc.EntityID, airStruc.EntityVal, airStruc.Message)

	updMessage, _ := json.Marshal(airStruc)
	// Write the state back to the ledger
	err = stub.PutState(A, updMessage)
	if err != nil {
		return shim.Error(err.Error())
	}

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
	var A string // Entities
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	A = args[0]

	// Get the state from the ledger
	Avalbytes, err := stub.GetState(A)

	airStruc := AirforceComm{}
	json.Unmarshal(Avalbytes, &airStruc)

	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + airStruc.Message + airStruc.Message + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return shim.Success(Avalbytes)
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
