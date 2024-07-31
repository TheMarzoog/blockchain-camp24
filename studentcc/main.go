package main

import (
	"fmt"
	"studentcc/internal/handler"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func mainChaincode() {
	chaincode, err := contractapi.NewChaincode(new(handler.StudentHandler))

	if err != nil {
		fmt.Printf("error creating student chaincode: %v", err)
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("error starting student chaincode: %v", err)
	}
}

func main() {

	mainChaincode()
	// handler := handler.StudentHandler{
	// 	RepoContract: *memory.New(),
	// }

}
