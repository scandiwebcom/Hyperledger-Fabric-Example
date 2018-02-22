package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the workers structure, with 3 properties.  Structure tags are used by encoding/json library
type Worker struct {
	Name     string `json:"name"`
	Status   string `json:"status"`
	Position string `json:"position"`
}

/*
 * The Init method is called when the Smart Contract "workers" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "workers"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "queryWorker" {
		return s.queryWorker(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "addWorker" {
		return s.addWorker(APIstub, args)
	} else if function == "queryAllWorkers" {
		return s.queryAllWorkers(APIstub)
	} else if function == "changeWorkerPosition" {
		return s.changeWorkerPosition(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) queryWorker(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	workerAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(workerAsBytes)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	workers := []Worker{
		Worker{Name: "Valerijs", Status: "working", Position: "Team Lead"},
		Worker{Name: "Maks", Status: "working", Position: "Developer"},
	}

	i := 0
	for i < len(workers) {
		fmt.Println("i is ", i)
		workerAsBytes, _ := json.Marshal(workers[i])
		APIstub.PutState("WORKER"+strconv.Itoa(i), workerAsBytes)
		fmt.Println("Added", workers[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) addWorker(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	var worker = Worker{Name: args[1], Status: args[2], Position: args[3]}

	workerAsBytes, _ := json.Marshal(worker)
	APIstub.PutState(args[0], workerAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryAllWorkers(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "WORKER0"
	endKey := "WORKER999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllWorkers:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) changeWorkerPosition(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	workerAsBytes, _ := APIstub.GetState(args[0])
	worker := Worker{}

	json.Unmarshal(workerAsBytes, &worker)
	worker.Position = args[1]

	workerAsBytes, _ = json.Marshal(worker)
	APIstub.PutState(args[0], workerAsBytes)

	return shim.Success(nil)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
