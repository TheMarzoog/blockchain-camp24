package contract

import (
	"encoding/json"
	"fmt"
	"studentcc/internal/model/student"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type RepoContract struct {
	contractapi.Contract
}

func (r *RepoContract) GetStudent(ctx contractapi.TransactionContextInterface, ID string) (*student.Student, error) {
	studentJSON, err := ctx.GetStub().GetState(ID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: error %v", err)
	}
	if studentJSON == nil {
		return nil, fmt.Errorf("student %s does not exist", ID)
	}

	var student student.Student
	err = json.Unmarshal(studentJSON, &student)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (r *RepoContract) GetAllStudents(ctx contractapi.TransactionContextInterface) ([]*student.Student, error) {
	studentIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, fmt.Errorf("failed to read form world state: error %v", err)
	}
	defer studentIterator.Close()

	var students []*student.Student

	for studentIterator.HasNext() {
		result, err := studentIterator.Next()
		if err != nil {
			return nil, err
		}
		var student student.Student
		err = json.Unmarshal(result.Value, &student)
		if err != nil {
			return nil, err
		}

		students = append(students, &student)
	}

	return students, nil
}

func (r *RepoContract) PutStudent(ctx contractapi.TransactionContextInterface, student *student.Student) error {
	studentJSON, err := json.Marshal(student)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(student.ID, studentJSON)
}
