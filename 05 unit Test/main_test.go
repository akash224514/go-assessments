package math

import (
	"testing"
)

type addTest struct { // Declaring a struct for Additon function check
	arg1, arg2, expected int
}

type subtractTest struct { // Declaring a struct for Subtraction function check
	arg1, arg2, expected int
}

var addTests = []addTest{ //Declaring a slice with two arguments and expected result vairable for Add function
	addTest{1, 2, 3},
	addTest{3, 4, 7},
	addTest{5, 6, 11},
}

var subtractTests = []subtractTest{ //Declaring a slice with two arguments and expected result vairable for Subtract function
	subtractTest{2, 1, 1},
	subtractTest{4, 3, 1},
	subtractTest{6, 5, 1},
}

func TestSubtract(t *testing.T) {
	for _, test := range subtractTests {
		if output := Subtract(test.arg1, test.arg2); output != test.expected { // Passing arguments to our subtract function and testing an output to the expected output
			t.Errorf("Output %d is not equal to expected %d", output, test.expected)
		}
	}
}

func TestAdd(t *testing.T) {

	for _, test := range addTests {
		if output := Add(test.arg1, test.arg2); output != test.expected { // Passing arguments to our Add function and testing an output to the expected output
			t.Errorf("Output %d is not equal to expected %d", output, test.expected)
		}
	}
}
