package utils

import "testing"

func TestExampleFunction(t *testing.T) {
	result := ExampleFunction()
	expected := "expected result"
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}