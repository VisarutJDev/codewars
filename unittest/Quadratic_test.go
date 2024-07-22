package unittest

import (
	"fmt"
	"testing"

	"github.com/VisarutJDev/codewars/eightkyu"
)

func TestQuadratic(t *testing.T) {
	testsample := [][]int{
		{0, 1},
		{4, 9},
		{2, 6},
		{-5, -4},
	}
	testresult := [][]int{
		{1, -1, 0},
		{1, -13, 36},
		{1, -8, 12},
		{1, 9, 20},
	}
	for i := range testsample {
		test := eightkyu.Quadratic(testsample[i][0], testsample[i][1])
		fmt.Println(test)
		if test[0] != testresult[i][0] && test[1] != testresult[i][1] && test[2] != testresult[i][2] {
			t.Errorf(`Expected (%v, %v, %v)`, testresult[i][0], testresult[i][1], testresult[i][2])
		}
	}
}
