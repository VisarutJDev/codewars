package unittest

import (
	"fmt"
	"testing"

	"github.com/VisarutJDev/codewars/eightkyu"
)

func TestCountSheeps(t *testing.T) {
	testsample := []bool{true, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, true, true,
		false, false, true, true}
	testresult := 17
	test := eightkyu.CountSheeps(testsample)
	fmt.Println(test)
	if test != testresult {
		t.Errorf(`Expected (%v)`, testresult)
	}
}
