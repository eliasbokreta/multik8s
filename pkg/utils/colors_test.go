package utils

import (
	"testing"
)

func TestGetColorsFn(t *testing.T) {
	c1, c2 := GetColorsFn(0)

	if c1("test") != "test" {
		t.Fatal("Should not be different")
	}

	if c2("test") != "test" {
		t.Fatal("Should not be different")
	}
}

func TestGetColorFn(t *testing.T) {
	c1 := GetColorFn(0)

	if c1("test") != "test" {
		t.Fatal("Should not be different")
	}
}
