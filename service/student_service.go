package service

import (
	"errors"
	"go-simple-test/domain"
	"go-simple-test/repository"
)

type StudentService interface {
	RegisterNewStudent(payload domain.Student) error
	FindById(id int) (domain.Student, error)
}

type studentService struct {
	repo repository.StudentRepository
}

func (s *studentService) RegisterNewStudent(payload domain.Student) error {
	if payload.Name == "" {
		return errors.New("name is required")
	}
	return s.repo.Create(payload)
}

func (s *studentService) FindById(id int) (domain.Student, error) {
	return s.repo.Get(id)
}

func NewStudentService(repo repository.StudentRepository) StudentService {
	return &studentService{repo: repo}
}
