package helper

import (
	"fmt"
	"testing"
)

func TestGreetingError(t *testing.T) {
	var value = "fajra"
	var expectedResult = fmt.Sprintf("Hello %s", value)
	result := Greeting(value)
	if result != expectedResult {
		// unit test is failed
		t.Error(fmt.Sprintf("'%s' is not equal to '%s'", result, expectedResult))
	}
	// it would run
	println("Application has completed")
}

func TestGreetingFail(t *testing.T) {
	var value = "fajra"
	var expectedResult = fmt.Sprintf("Hello %s", value)
	result := Greeting(value)
	if result != expectedResult {
		// unit test is failed
		t.Fail()
	}
	// it would run
	println("Application has completed")
}

func TestGreetingFatal(t *testing.T) {
	var value = "fajra"
	var expectedResult = fmt.Sprintf("Hello %s", value)
	result := Greeting(value)
	if result != expectedResult {
		// unit test is failed
		t.Fatal(fmt.Sprintf("'%s' is not equal to '%s'", result, expectedResult))
	}
	// it would not run
	println("Application has completed")
}

func TestGreetingFailNow(t *testing.T) {
	var value = "fajra"
	var expectedResult = fmt.Sprintf("Hello %s", value)
	result := Greeting(value)
	if result != expectedResult {
		// unit test is failed
		t.FailNow()
	}
	// it would not run
	println("Application has completed")
}

func TestGreeting(t *testing.T) {
	var value = "fajra"
	var expectedResult = fmt.Sprintf("Hello %s", value)
	result := Greeting(value)
	if result != expectedResult {
		// unit test is failed
		panic(fmt.Sprintf("'%s' is not equal to '%s'", result, expectedResult))
	}
}

func TestGreetingCisnux(t *testing.T) {
	var value = "cisnux"
	var expectedResult = fmt.Sprintf("Hello %s", value)
	result := Greeting(value)
	if result != expectedResult {
		// unit test is failed
		panic(fmt.Sprintf("'%s' is not equal to '%s'", result, expectedResult))
	}
}
