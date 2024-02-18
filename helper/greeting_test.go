package helper

import (
	"fmt"
	"testing"
)

func TestGreeting(t *testing.T) {
	var value = "fajra"
	var expectedResult = fmt.Sprintf("Hi %s", value)
	result := Greeting(value)
	if result != expectedResult {
		// unit test is failed
		panic(fmt.Sprintf("%s is not equal to %s", result, expectedResult))
	}
}

func TestGreetingCisnux(t *testing.T) {
	var value = "cisnux"
	var expectedResult = fmt.Sprintf("Hi %s", value)
	result := Greeting(value)
	if result != expectedResult {
		// unit test is failed
		panic(fmt.Sprintf("%s is not equal to %s", result, expectedResult))
	}
}
