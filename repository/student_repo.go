package repository

import (
	"errors"
	"fmt"
	"go-simple-test/domain"
)

type StudentRepository interface {
	Create(payload domain.Student) error
	Get(id int) (domain.Student, error)
}

type studentRepository struct {
	db []domain.Student
}

func (s *studentRepository) Create(payload domain.Student) error {
	if payload.Name == "" {
		return errors.New("name is required")
	}
	s.db = append(s.db, payload)
	return nil
}

func (s *studentRepository) Get(id int) (domain.Student, error) {
	var student domain.Student
	for _, v := range s.db {
		if v.Id != id {
			return domain.Student{}, fmt.Errorf("student with ID %v not found", id)
		}

		student = v
	}
	return student, nil
}

func NewStudentRepository() StudentRepository {
	studentDummies := []domain.Student{
		{
			Id:   1,
			Name: "Suzuki",
		},
		{
			Id:   2,
			Name: "Honda",
		},
	}
	return &studentRepository{
		db: studentDummies,
	}
}
