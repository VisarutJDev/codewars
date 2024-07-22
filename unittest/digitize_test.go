package unittest

import (
	"reflect"
	"testing"

	"github.com/VisarutJDev/codewars/eightkyu"
)

func TestDigitize(t *testing.T) {
	// your code here
	// 	35231 => [1,3,2,5,3]
	// 0 => [0]
	testsample := []int{
		35231,
		0,
	}
	testResult := [][]int{
		{1, 3, 2, 5, 3},
		{0},
	}
	for i := range testsample {
		result := eightkyu.Digitize(testsample[i])
		if !reflect.DeepEqual(result, testResult[i]) {
			t.Errorf(`Expected (%v)`, testResult[i])
		}
	}

}
