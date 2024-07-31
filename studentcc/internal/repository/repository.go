package repository

import (
	"studentcc/internal/model/student"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Repository interface {
	GetStudent(contractapi.TransactionContextInterface, string) (*student.Student, error)
	GetAllStudents(contractapi.TransactionContextInterface) ([]*student.Student, error)
	PutStudent(contractapi.TransactionContextInterface, *student.Student) error
}
