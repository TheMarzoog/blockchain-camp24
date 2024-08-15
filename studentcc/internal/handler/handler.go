package handler

import (
	"fmt"
	"studentcc/internal/model/student"

	"studentcc/internal/repository/memory"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type StudentHandler struct {
	RepoContract memory.RepoMemory
	// contract.RepoContract
}

func (h *StudentHandler) GetStudnet(ctx contractapi.TransactionContextInterface, ID string) (*student.Student, error) {
	return h.RepoContract.GetStudent(ctx, ID)
}

func (h *StudentHandler) GetAllStudents(ctx contractapi.TransactionContextInterface) ([]*student.Student, error) {
	return h.RepoContract.GetAllStudents(ctx)
}

func (h *StudentHandler) Register(ctx contractapi.TransactionContextInterface, ID, name, email, phone, camp string) error {
	// Check if the student is Registered
	_, err := h.GetStudnet(ctx, ID)
	if err == nil {
		return fmt.Errorf("studnent with ID %s is registered", ID)
	}
	student := student.New(ID, name, email, phone, camp)
	return h.RepoContract.PutStudent(ctx, student)
}

func (h *StudentHandler) UpdateEmail(ctx contractapi.TransactionContextInterface, ID, newEmail string) error {
	// check if the student exists
	student, err := h.GetStudnet(ctx, ID)
	if err != nil {
		return err
	}
	student.Email = newEmail
	return h.RepoContract.PutStudent(ctx, student)
}

func (h *StudentHandler) UpdatePhone(ctx contractapi.TransactionContextInterface, ID, newPhone string) error {
	// check if the student exists
	student, err := h.GetStudnet(ctx, ID)
	if err != nil {
		return err
	}
	student.Phone = newPhone
	return h.RepoContract.PutStudent(ctx, student)
}

func (h *StudentHandler) ChangeCamp(ctx contractapi.TransactionContextInterface, ID, newCamp string) error {
	// check if the student exists
	student, err := h.GetStudnet(ctx, ID)
	if err != nil {
		return err
	}
	student.Camp = newCamp
	return h.RepoContract.PutStudent(ctx, student)
}

func (h *StudentHandler) Grade(ctx contractapi.TransactionContextInterface, ID string, grade float64) error {
	if grade < 0 || grade > 100 {
		return fmt.Errorf("%.2f is not between 0 and 100", grade)
	}
	// check if the student exists
	student, err := h.GetStudnet(ctx, ID)
	if err != nil {
		return err
	}
	student.ProjectGrade = grade
	return h.RepoContract.PutStudent(ctx, student)
}
