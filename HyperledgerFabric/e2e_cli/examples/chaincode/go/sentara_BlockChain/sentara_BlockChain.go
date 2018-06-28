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
	//"github.com/back_hyp/hyperledger/blockchain/trunk/0.6/fabric/membersrvc/protos"
	//"github.com/hyperledger/fabric/gossip/election"
)

var start time.Time
var countMess int

//custom data models
type sentara struct {
	EntityID string `json:"entityID"`
	EntityVal string `json:"entityVal"`
	Message string `json:"message"`
	Intent string `json:"intent"`
	Port string `json:"port"`
	Packet string `json:"packet"`

}

var m map[string][]byte

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

func createMap() bool {
	for i := 0; i < 100; i++ {
		theObj := sentara{}
		theObj.EntityID = "a"+strconv.Itoa(i)
		theObj.EntityVal = "bank"
		theObj.Message = "test"
		theObj.Intent = "http"
		theObj.Port = "80"
		theObj.Packet = "256"
		marshMess, _ := json.Marshal(theObj)
		m["a"+strconv.Itoa(i)] = marshMess

	}
	return true
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Init")
	_, args := stub.GetFunctionAndParameters()
	var A, B, mess string    // Entities
	//var messagStr = []byte
	//var Aval, Bval int // Asset holdings
	var err error

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	m = make(map[string][]byte)
	createMap()

	// Initialize the chaincode
	A = args[0]
	B = args[1]

	mess = args[2]
	fmt.Printf("Aval = %d, Bval = %d\n, mess = %s/n", A, B, mess)

	// Write the state to the ledger
	//messagStr = "{\"entityID\":" + A + ",\"entityVal\":" + B + ",\"message\":"+mess+"}"
	messagStr, _ := json.Marshal(m)

	err = stub.PutState("A", messagStr)
	if err != nil {
		return shim.Error(err.Error())
	}

	//messagStr = "{\"entityID\":" + B + ",\"entityVal\":" + A + ",\"message\":"+mess+"}"
	//err = stub.PutState("B", messagStr)
	//if err != nil {
    //	return shim.Error(err.Error())
	//}

	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Invoke")
	function, args := stub.GetFunctionAndParameters()
	if function == "invoke" {
		// invokes the item to the blockchain (must be a sentara object)
		return t.invoke(stub, args)
	} else if function == "invokeBatch" {
		// provides invoke service to local data structure (not blockchain)
		return t.invokeBatch(stub, args)
	} else if function == "delete" {
		// todo currently just deletes item on blockchain
		return t.delete(stub, args)
	} else if function == "query" {
		// Individual query, checks for id on the blockchain
		return t.query(stub, args)
	} else if function == "add" {
		return t.add(stub, args)
	} else if function == "batch" {
		return t.batch(stub, args)
	} else if function == "auto" {
		return t.autoVoke(stub, args)
	}

	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}

//grab value, unmarshal value into sentara obj
//update the message
//marshal the object and save as new value to existing key
//if so then update the message
func mUpdater(a string, b string, msg string, intnt string, prt string, pckt string) sentara {
	var afc = sentara {}
	json.Unmarshal(m[a], &afc)
	afc.EntityID=a
	afc.EntityVal=b
	afc.Message=msg
	afc.Intent=intnt
	afc.Port=prt
	afc.Packet=pckt
	marshMess, _ := json.Marshal(afc)
	m[a] = marshMess
	return afc
}

//Adds a new ID entity to the blockchain
func (t *SimpleChaincode) add(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var A, B, mess, messagStr, intnt, prt, pckt string    // Entities
	var err error

	if len(args) == 3{
		A = args[0]
		B = args[1]
		mess = args[2]
		intnt = "http"
		prt = "80"
		pckt = "256"
	} else if len(args) == 6 {
		A = args[0]
		B = args[1]
		mess = args[2]
		intnt = args[3]
		prt = args[4]
		pckt = args[5]
	} else {
		return shim.Error("Incorrect number of arguments. Expecting 3 or 6")
	}

	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger

	messagStr = "{\"entityID\":" + A + ",\"entityVal\":" + B + ",\"message\":"+mess+"}"
	fmt.Printf(messagStr)

	if _, ok := m[A]; ok {
		mUpdater(A, B, mess, intnt, prt, pckt)
	} else {
		nAFC := sentara{}
		nAFC.EntityID = A
		nAFC.EntityVal = B
		nAFC.Message = mess
		nAFC.Intent = "http"
		nAFC.Port = "80"
		nAFC.Packet = "256"
		marshMess, _ := json.Marshal(nAFC)
		m[A] = marshMess
	}

	marshMess, _ := json.Marshal(m)
	err = stub.PutState("A", marshMess)
	if err != nil {
		return shim.Error("error with A PutState")
	}

	return shim.Success(nil)
}

func (t *SimpleChaincode) addBatch(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var A, B, mess, messagStr, intnt, prt, pckt string    // Entities
	var err error

	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	A = args[0]
	B = args[1]
	mess = args[2]
	intnt = args[3]
	prt = args[4]
	pckt = args[5]

	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger

	messagStr = "{\"entityID\":" + A + ",\"entityVal\":" + B + ",\"message\":"+mess+"}"
	fmt.Printf(messagStr)

	if _, ok := m[A]; ok {
		mUpdater(A, B, mess, intnt, prt, pckt)
	} else {
		nAFC := sentara{}
		nAFC.EntityID = A
		nAFC.EntityVal = B
		nAFC.Message = mess
		nAFC.Intent = "http"
		nAFC.Port = "80"
		nAFC.Packet = "256"
		marshMess, _ := json.Marshal(nAFC)
		m[A] = marshMess
	}

	marshMess, _ := json.Marshal(m)
	err = stub.PutState("A", marshMess)
	if err != nil {
		return shim.Error("error with A PutState")
	}

	return shim.Success(nil)
}

// Create a batch set test for key value par id'ed a
// the batch set contains objects for values a1-a10000
// this allows for local and permanent block storage
func (t *SimpleChaincode) batch(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var A string    // Entities
	//var Aval, Bval int // Asset holdings
	//var X int          // Transaction value
	var err error
	A = args[0]
    //id1 := [5]string{"c","e","g","i","k"}
    //id2 := [5]string{"d","f","h","j","l"}

	for q := 0; q < 100; q++ {
		if _, ok := m["a"+strconv.Itoa(q)]; ok {
			mUpdater("a"+strconv.Itoa(q), "bank", "batcher", "http", "80", "256")
		} else {
			nAFC := sentara{}
			nAFC.EntityID = "a"+strconv.Itoa(q)
			nAFC.EntityVal = "bank"
			nAFC.Message = "batcher"
			nAFC.Intent = "http"
			nAFC.Port = "80"
			nAFC.Packet = "256"
			marshMess, _ := json.Marshal(nAFC)
			m["a"+strconv.Itoa(q)] = marshMess
		}
	}

	updMessage, _ := json.Marshal(m)
		//fmt.Printf("Aval = %s, Bval = %s, mess = %s\n", A, B, mess)
		// Write the state back to the ledger
	err = stub.PutState("A", updMessage)
	if err != nil {
		return shim.Error(err.Error())
	} else {
		fmt.Printf("key %s does not exist... Need to use the Add function", A)
	}

	return shim.Success(nil)
}

// This function only submits the test batch object to the blockchain
// the idea for this function is to periodically write objects to the blockchain
// this writes the object as a state on the blockchain to make it permanenet and immutable
func (t *SimpleChaincode) autoVoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	updMessage, _ := json.Marshal(m)
	err = stub.PutState("A", updMessage)
		if err != nil {
			return shim.Error(err.Error())
		}
	fmt.Printf("key %s does not exist... Need to use the Add function", "a")

	return shim.Success(nil)
}

//this function updates the data structure locally not permanently on the blockchain
//The chaincode relies on the autoVoke function in order to permanently write the data structure to the blockchain
func (t *SimpleChaincode) invokeBatch(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var A, B, mess, intnt, prt, pckt string    // Entities
	//var Aval, Bval int // Asset holdings
	//var X int          // Transaction value
	var err error

	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	A = args[0]
	B = args[1]
	mess = args[2]
	intnt = args[3]
	prt = args[4]
	pckt = args[5]

	if _, ok := m[A]; ok {
		//mUpdater(A, B, mess)
		mUpdater(A, B, mess, intnt, prt, pckt)
		//updMessage, _ := json.Marshal(m)
		fmt.Printf("Batch Aval = %s, Bval = %s, mess = %s\n", A, B, mess)
		// Write the state back to the ledger
		//err = stub.PutState("A", updMessage)
		if err != nil {
			return shim.Error(err.Error())
		}
	} else {
		fmt.Printf("key %s does not exist... Need to use the Add function", A)
	}

	return shim.Success(nil)
}

// Makes an independent transaction not associated with a batch id
// The input is the id for the independent entity that is being transacted on
func (t *SimpleChaincode) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var A, B, mess, intnt, prt, pckt string    // Entities
	//var Aval, Bval int // Asset holdings
	//var X int          // Transaction value
	var err error

	if len(args) == 3 {
		A = args[0]
		B = args[1]
		mess = args[2]
		intnt = "http"
		prt = "80"
		pckt = "256"
	} else if len(args) == 6 {
		A = args[0]
		B = args[1]
		mess = args[2]
		intnt = args[3]
		prt = args[4]
		pckt = args[5]
	} else {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	if _, ok := m[A]; ok {
		mUpdater(A, B, mess, intnt, prt, pckt)
	} else {
		nAFC := sentara{}
		nAFC.EntityID = A
		nAFC.EntityVal = B
		nAFC.Message = mess
		nAFC.Intent = "http"
		nAFC.Port = "80"
		nAFC.Packet = "256"
		marshMess, _ := json.Marshal(nAFC)
		m[A] = marshMess
	}

	marshMess, _ := json.Marshal(m)
	err = stub.PutState("A", marshMess)
	if err != nil {
		return shim.Error("error with A PutState")
	}

	return shim.Success(nil)
}

// ToDo check if id exists in blockchain or in local data structure
// ToDo then remove the item from the blockchain or data structure
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

//used to query data that is stored in batch invoke (locally but checked with blockchain)
func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var A string // Entities
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	A = args[0]

	// Get the state from the ledger
	Avalbytes, err := stub.GetState("A")

	tempM := make(map[string][]byte)
	sentaraVal := sentara{}

	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		return shim.Error(jsonResp)
	}

    if marshErr := json.Unmarshal(Avalbytes, &tempM); marshErr != nil {
		checkBytes := Avalbytes
    	if singErr := json.Unmarshal(Avalbytes, &sentaraVal); singErr != nil {
			return shim.Error("Data structure is not recognized")
		}
		jsonResp := "{\"SingleName\":\"" + A + "\",\"Amount\":\"" + sentaraVal.Message +" "+ sentaraVal.Message + "\"}"
		fmt.Printf("Query Response:%s\n", jsonResp)

		return shim.Success(checkBytes)

	} else {
		nvalbytes := tempM[A]
		json.Unmarshal(nvalbytes, &sentaraVal)

		checkBytes := m[A]
		checkStruc := sentara{}
		json.Unmarshal(checkBytes, &checkStruc)
		jsonResp := "{\"BatchName\":\"" + A + "\",\"Amount\":\"" + checkStruc.Message +" "+ sentaraVal.Message + "\"}"
		fmt.Printf("Query Response:%s\n", jsonResp)

		return shim.Success(nvalbytes)
	}

}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
