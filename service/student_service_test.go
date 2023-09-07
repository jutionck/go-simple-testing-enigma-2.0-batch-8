package service

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go-simple-test/domain"
	"testing"
)

type studentRepoMock struct {
	mock.Mock
}

func (s *studentRepoMock) Create(payload domain.Student) error {
	args := s.Called(payload)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (s *studentRepoMock) Get(id int) (domain.Student, error) {
	args := s.Called(id)
	if args.Get(1) != nil {
		return domain.Student{}, args.Error(1)
	}
	return args.Get(0).(domain.Student), nil
}

type StudentServiceTestSuite struct {
	suite.Suite
	repoMock *studentRepoMock
	service  StudentService
}

func (suite *StudentServiceTestSuite) SetupTest() {
	suite.repoMock = new(studentRepoMock)
	suite.service = NewStudentService(suite.repoMock)
}

func TestStudentServiceTestSuite(t *testing.T) {
	suite.Run(t, new(StudentServiceTestSuite))
}

func (suite *StudentServiceTestSuite) TestRegisterNewStudent_Success() {
	payload := domain.Student{
		Id:   1,
		Name: "Suzuki",
	}
	suite.repoMock.On("Create", payload).Return(nil)
	err := suite.service.RegisterNewStudent(payload)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *StudentServiceTestSuite) TestRegisterNewStudent_Fail() {
	payload := domain.Student{
		Id:   1,
		Name: "",
	}
	expectedError := errors.New("name is required")
	suite.repoMock.On("Create", payload).Return(expectedError)
	err := suite.service.RegisterNewStudent(payload)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), err.Error(), expectedError.Error())
}

func (suite *StudentServiceTestSuite) TestFindById_Success() {
	expected := domain.Student{
		Id:   1,
		Name: "Suzuki",
	}
	suite.repoMock.On("Get", expected.Id).Return(expected, nil)
	got, err := suite.service.FindById(expected.Id)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, got)
}

func (suite *StudentServiceTestSuite) TestFindById_Fail() {
	mockStudentID := 0
	expectedError := fmt.Errorf("student with ID %v not found", mockStudentID)
	suite.repoMock.On("Get", mockStudentID).Return(domain.Student{}, expectedError)
	_, err := suite.service.FindById(mockStudentID)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedError.Error(), err.Error())
}
