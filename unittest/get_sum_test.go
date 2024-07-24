package unittest

import (
	"reflect"
	"testing"

	"github.com/VisarutJDev/codewars/sevenkyu"
)

func TestGetSum(t *testing.T) {
	testsample := [][]int{
		{0, 1},
		{1, 2},
		{5, -1},
		{505, 4},
		{321, 123},
		{0, -1},
		{-50, 0},
		{-1, -5},
		{-5, -5},
		{-505, 4},
		{-321, 123},
		{0, 0},
		{-5, -1},
		{5, 1},
		{-17, -17},
		{17, 17},
	}
	testResult := []int{
		1,
		3,
		14,
		127759,
		44178,
		-1,
		-1275,
		-15,
		-5,
		-127755,
		-44055,
		0,
		-15,
		15,
		-17,
		17,
	}
	for i := range testsample {
		result := sevenkyu.GetSum(testsample[i][0], testsample[i][1])
		if !reflect.DeepEqual(result, testResult[i]) {
			t.Errorf(`Expected (%v)`, testResult[i])
		}
	}

}
