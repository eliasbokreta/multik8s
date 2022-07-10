package utils

import (
	"testing"
)

func TestAgeFormatter(t *testing.T) {
	value := uint64(86400)
	age := AgeFormatter(value)
	if age != "1d" {
		t.Fatalf("Age should be equal to 1d: input: %d - output: %s", value, age)
	}

	value = uint64(3600)
	age = AgeFormatter(value)
	if age != "1h" {
		t.Fatalf("Age should be equal to 1h: %d - output: %s", value, age)
	}

	value = uint64(60)
	age = AgeFormatter(value)
	if age != "1m" {
		t.Fatalf("Age should be equal to 1m: %d - output: %s", value, age)
	}

	value = uint64(1)
	age = AgeFormatter(value)
	if age != "1s" {
		t.Fatalf("Age should be equal to 1s: %d - output: %s", value, age)
	}

	value = uint64(0)
	age = AgeFormatter(value)
	if age != "0s" {
		t.Fatalf("Age should be equal to 0s: %d - output: %s", value, age)
	}

	value = uint64(2000)
	age = AgeFormatter(value)
	if age != "33m20s" {
		t.Fatalf("Age should be equal to 33m20s: %d - output: %s", value, age)
	}

	value = uint64(4000)
	age = AgeFormatter(value)
	if age != "1h6m" {
		t.Fatalf("Age should be equal to 1h6m: %d - output: %s", value, age)
	}
}
