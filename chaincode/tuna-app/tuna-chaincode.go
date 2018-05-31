// SPDX-License-Identifier: Apache-2.0

/*
  Sample Chaincode based on Demonstrated Scenario

 This code is based on code written by the Hyperledger Fabric community.
  Original code can be found here: https://github.com/hyperledger/fabric-samples/blob/release/chaincode/fabcar/fabcar.go
 */

package main

/* Imports  
* 4 utility libraries for handling bytes, reading and writing JSON, 
formatting, and string manipulation  
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

/* Define Tuna structure, with 4 properties.  
Structure tags are used by encoding/json library
*/
// type Tuna struct {
// 	Vessel string `json:"vessel"`
// 	Timestamp string `json:"timestamp"`
// 	Location  string `json:"location"`
// 	Holder  string `json:"holder"`
//}

type Medicament struct {
    RefProduit string `json:"refProduit"`
	NomProduit string `json:"nomProduit"`
	NomFabricant  string `json:"nomFabricant"`
    NumeroLot string `json:"numeroLot"`
	DateExp string `json:"dateExp"`
	LocalisationVille string `json:"localisationVille"`
	LocalisationEtablissement string `json:"localisationEtablissement"`
	DateCreation string `json:"dateCreation"`
	DateReception string `json:"dateReception"`
	VenteClient string `json:"venteClient"`
}


/*
 * The Init method *
 called when the Smart Contract "tuna-chaincode" is instantiated by the network
 * Best practice is to have any Ledger initialization in separate function 
 -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method *
 called when an application requests to run the Smart Contract "tuna-chaincode"
 The app also specifies the specific smart contract function to call with args
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()

	fmt.Println(function)

	// Route to the appropriate handler function to interact with the ledger
	if function == "queryMedic" {
		return s.queryMedic(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "recordMedic" {
		return s.recordMedic(APIstub, args)
	} else if function == "queryAllMedic" {
		return s.queryAllMedic(APIstub)
	}

	return shim.Error("COINCOUN.")
}

/*
 * The queryTuna method *
Used to view the records of one particular tuna
It takes one argument -- the key for the tuna in question
 */
func (s *SmartContract) queryMedic(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	medicAsBytes, _ := APIstub.GetState(args[0])
	if medicAsBytes == nil {
		return shim.Error("Could not locate medic")
	}
	return shim.Success(medicAsBytes)
}

/*
 * 	Ici, on initialise notre JDD sur le network
 */
func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {

	medic := []Medicament{
		Medicament{RefProduit: "", NomProduit:"Doliprane", NomFabricant: "Sanofi", NumeroLot: "ABCDEF1234567", DateExp: "16/02/2020", 
		LocalisationVille: "Paris", LocalisationEtablissement: "Etablissement1", DateCreation:"06/08/2017", DateReception: "", VenteClient: "true"},

		Medicament{RefProduit: "", NomProduit:"Doliprane", NomFabricant: "Sanofi", NumeroLot: "ABCDEF1234567", DateExp: "16/02/2020", 
		LocalisationVille: "Talence", LocalisationEtablissement: "Etablissement2", DateCreation:"06/08/2017", DateReception: "", VenteClient: "true"},

		Medicament{RefProduit: "", NomProduit:"Doliprane", NomFabricant: "Sanofi", NumeroLot: "ABCDEF1234567", DateExp: "16/02/2020", 
		LocalisationVille: "Bordeaux", LocalisationEtablissement: "Etablissement4", DateCreation:"06/08/2017", DateReception: "", VenteClient: "true"},

		Medicament{RefProduit: "", NomProduit:"Exomuc", NomFabricant: "Sanofi", NumeroLot: "AA1235458AZEE", DateExp: "06/05/2022", 
		LocalisationVille: "Bordeaux", LocalisationEtablissement: "Etablissement1", DateCreation:"", DateReception: "", VenteClient: "true"},
	}

	i := 0
	for i < len(medic) {
		fmt.Println("i is ", i)
		medicAsBytes, _ := json.Marshal(medic[i])
		APIstub.PutState(strconv.Itoa(i+1), medicAsBytes)
		fmt.Println("Added", medic[i])
		i = i + 1
	}

	return shim.Success(nil)
}

/*
 * The recordTuna method *
Fisherman like Sarah would use to record each of her tuna catches. 
This method takes in five arguments (attributes to be saved in the ledger). 
 */
func (s *SmartContract) recordMedic(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 11 {
		return shim.Error("Incorrect number of arguments. Expecting 10")
	}

	var medic = Medicament{ RefProduit: args[1], NomProduit: args[2], NomFabricant: args[3],
	 NumeroLot: args[4], DateExp: args[5], LocalisationVille: args[6], LocalisationEtablissement: args[7], DateCreation: args[8], DateReception: args[9], VenteClient: args[10]}

	medicAsBytes, _ := json.Marshal(medic)
	err := APIstub.PutState(args[0], medicAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record medic catch: %s", args[0]))
	}

	return shim.Success(nil)
}

/*
 * The queryAllTuna method *
allows for assessing all the records added to the ledger(all tuna catches)
This method does not take any arguments. Returns JSON string containing results. 
 */
func (s *SmartContract) queryAllMedic(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "0"
	endKey := "999"

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
		// Add comma before array members,suppress it for the first array member
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

	fmt.Printf("- queryAllMedic:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}


/*
 * main function *
calls the Start function 
The main function starts the chaincode in the container during instantiation.
 */
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
