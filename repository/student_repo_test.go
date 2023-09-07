package repository

import (
	"errors"
	"fmt"
	"go-simple-test/domain"
	"testing"
)

func TestStudentServiceCreate_Success(t *testing.T) {
	stdRepo := NewStudentRepository()
	mock := domain.Student{
		Id:   1,
		Name: "Suzuki",
	}
	err := stdRepo.Create(mock)
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}
}

func TestStudentServiceCreate_Fail(t *testing.T) {
	stdRepo := NewStudentRepository()
	mock := domain.Student{
		Id:   1,
		Name: "",
	}
	err := stdRepo.Create(mock)
	expected := errors.New("name is required")
	if err != nil {
		if err.Error() != expected.Error() {
			t.Errorf("expected error: %v, but got: %v", err.Error(), expected.Error())
		}
	}
}

func TestStudentServiceGet_Success(t *testing.T) {
	stdRepo := NewStudentRepository()
	fmt.Println(stdRepo.Get(3))
	expectedStudent := domain.Student{Id: 1, Name: "Suzuki"}
	student, err := stdRepo.Get(1)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if student != expectedStudent {
		t.Errorf("Expected student: %v, but got: %v", expectedStudent, student)
	}
}

func TestStudentServiceGet_Fail(t *testing.T) {
	stdRepo := NewStudentRepository()
	mockStudentID := 0
	_, err := stdRepo.Get(0)
	expected := fmt.Errorf("student with ID %v not found", mockStudentID)
	if err != nil {
		if err.Error() != expected.Error() {
			t.Errorf("expected error: %v, but got: %v", err.Error(), expected.Error())
		}
	}
}
