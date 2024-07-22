package unittest

import (
	"testing"

	"github.com/VisarutJDev/codewars/sevenkyu"
)

func TestRowSumOddNumbers(t *testing.T) {
	// fill in solution here
	testsample := []int{
		1,
		2,
		13,
		19,
		41,
		42,
		74,
		86,
		93,
		101,
	}
	testresult := []int{
		1,
		8,
		2197,
		6859,
		68921,
		74088,
		405224,
		636056,
		804357,
		1030301,
	}
	for i := range testsample {
		test := sevenkyu.RowSumOddNumbers(testsample[i])
		if test != testresult[i] {
			t.Errorf(`Expected (%v) But got (%v)`, testresult[i], test)
		}
	}

}
