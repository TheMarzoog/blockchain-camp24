package memory

import (
	"fmt"
	"studentcc/internal/model/student"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type RepoMemory map[string]*student.Student

func New() *RepoMemory {
	rm := make(RepoMemory)
	return &rm
}

func (rm *RepoMemory) GetStudent(_ contractapi.TransactionContextInterface, ID string) (*student.Student, error) {
	student, existing := (*rm)[ID]
	if !existing {
		return nil, fmt.Errorf("no student with ID: %s", ID)
	}
	return student, nil
}

func (rm *RepoMemory) GetAllStudents(_ contractapi.TransactionContextInterface) ([]*student.Student, error) {
	var students []*student.Student

	for _, v := range *rm {
		students = append(students, v)
	}

	if len(students) == 0 {
		return nil, fmt.Errorf("no students found")
	}

	return students, nil
}

// used to add new student and to update
func (rm *RepoMemory) PutStudent(_ contractapi.TransactionContextInterface, student *student.Student) error {
	(*rm)[student.ID] = student
	return nil
}
