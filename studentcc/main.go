package main

import (
	"fmt"
	"studentcc/internal/handler"

	"studentcc/internal/repository/memory"
)

// func mainChaincode() {
// 	chaincode, err := contractapi.NewChaincode(new(handler.StudentHandler))

// 	if err != nil {
// 		fmt.Printf("error creating student chaincode: %v", err)
// 		return
// 	}

// 	if err := chaincode.Start(); err != nil {
// 		fmt.Printf("error starting student chaincode: %v", err)
// 	}
// }

func main() {

	// mainChaincode()
	handler := handler.StudentHandler{
		RepoContract: *memory.New(),
	}

	err := handler.Register(nil, "201229293", "Marzoog", "201229293@pus.edu.sa", "059198283", "Blockchain Camp")
	if err != nil {
		panic(err)
	}

	student, err := handler.GetStudnet(nil, "201229293")

	if err != nil {
		panic(err)
	}

	fmt.Printf("- %#v", student)

}
