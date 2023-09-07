package service

import (
	"errors"
	"testing"
)

func TestCalculatorAdd_Success(t *testing.T) {
	var cal = &Calculator{
		Number1: 6,
		Number2: 4,
	}

	var expected float64 = 10
	actual, _ := cal.Add()
	if *actual != expected {
		t.Errorf("expected result: %v, but got result: %v", expected, *actual)
	}
}

func TestCalculatorAdd_Fail(t *testing.T) {
	var cal = &Calculator{
		Number1: -6,
		Number2: 4,
	}
	var expectedError = errors.New("number can't be negative")
	_, actualError := cal.Add()
	if actualError != nil {
		if actualError.Error() != expectedError.Error() {
			t.Errorf("expected error: %v, but got error: %v", expectedError, actualError)
		}
	}
}

func TestCalculatorSub_Success(t *testing.T) {
	var cal = &Calculator{
		Number1: 6,
		Number2: 4,
	}

	var expected float64 = 2
	actual, _ := cal.Sub()
	if *actual != expected {
		t.Errorf("expected result: %v, but got result: %v", expected, *actual)
	}
}

func TestCalculatorSub_Fail(t *testing.T) {
	var cal = &Calculator{
		Number1: -6,
		Number2: 4,
	}
	var expectedError = errors.New("number can't be negative")
	_, actualError := cal.Sub()
	if actualError != nil {
		if actualError.Error() != expectedError.Error() {
			t.Errorf("expected error: %v, but got error: %v", expectedError, actualError)
		}
	}
}
