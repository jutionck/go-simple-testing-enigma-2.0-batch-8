package repository

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
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
	assert.Nil(t, err)
}

func TestStudentServiceCreate_Fail(t *testing.T) {
	stdRepo := NewStudentRepository()
	mock := domain.Student{
		Id:   1,
		Name: "",
	}
	err := stdRepo.Create(mock)
	expected := errors.New("name is required")
	assert.Error(t, err)
	assert.Equal(t, expected.Error(), err.Error())
}

func TestStudentServiceGet_Success(t *testing.T) {
	stdRepo := NewStudentRepository()
	expectedStudent := domain.Student{Id: 1, Name: "Suzuki"}
	student, err := stdRepo.Get(1)
	assert.Nil(t, err)
	assert.Equal(t, expectedStudent, student)
}

func TestStudentServiceGet_Fail(t *testing.T) {
	stdRepo := NewStudentRepository()
	mockStudentID := 0
	_, err := stdRepo.Get(0)
	expected := fmt.Errorf("student with ID %v not found", mockStudentID)
	assert.Error(t, err)
	assert.Equal(t, expected.Error(), err.Error())
}
