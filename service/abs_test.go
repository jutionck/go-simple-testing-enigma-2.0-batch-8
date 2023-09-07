package service

import (
	"math"
	"testing"
)

func TestAbs(t *testing.T) {
	got := math.Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %f; want 1", got)
	}
}

func TestOther(t *testing.T) {
	got := math.Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %f; want 1", got)
	}
}

// go test -> menjalankan unit test
// go test -v => menjalankan unit test secara detail
// go test -v -run namaFunction => menjalankan unit test secara detail dan spesifik function mana
// go test -v ./... => menjalankan unit test yang file testingnya terletak di berbagai package
