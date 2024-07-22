package unittest

import (
	"testing"

	"github.com/VisarutJDev/codewars/eightkyu"
)

// Make a function that will return a greeting statement that uses an input; your program should return, "Hello, <name> how are you doing today?".

// [Make sure you type the exact thing I wrote or the program may not execute properly]

// Returning Strings

func TestGreet(t *testing.T) {
	// fill in solution here
	testsample := `Ryan`
	testresult := `Hello, Ryan how are you doing today?`
	test := eightkyu.Greet(testsample)

	if test != testresult {
		t.Errorf(`Expected (%v) But got (%v)`, testresult, test)
	}
}
