package main

import (
	"example/process"
	"testing"
)

func TestCheckEven(t *testing.T) {
	i := 2
	expected := "YES"
	response := process.CheckEven(i)

	if expected != response {
		t.Errorf("Expected %v, got %v", expected, response)
	}

}
