package service

import "errors"

// Calculator -> untuk menghitung (+,-,/, *)

type Calculator struct {
	Number1 float64
	Number2 float64
}

func (c *Calculator) Add() (*float64, error) {
	if c.Number1 < 0 || c.Number2 < 0 {
		return nil, errors.New("number can't be negative")
	}
	result := c.Number1 + c.Number2
	return &result, nil
}

func (c *Calculator) Sub() (*float64, error) {
	if c.Number1 < 0 || c.Number2 < 0 {
		return nil, errors.New("number can't be negative")
	}
	result := c.Number1 - c.Number2
	return &result, nil
}
