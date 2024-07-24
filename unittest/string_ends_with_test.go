package unittest

import (
	"reflect"
	"testing"

	"github.com/VisarutJDev/codewars/sevenkyu"
)

func TestSolution(t *testing.T) {
	// Your code here!
	testsample := [][]string{
		{"", ""},
		{" ", ""},
		{"abc", "c"},
		{"banana", "ana"},
		{"a", "z"},
		{"", "t"},
		{"$a = $b + 1", "+1"},
		{"    ", "   "},
		{"1oo", "100"},
	}
	testResult := []bool{
		true,
		true,
		true,
		true,
		false,
		false,
		false,
		true,
		false,
	}
	for i := range testsample {
		result := sevenkyu.Solution(testsample[i][0], testsample[i][1])
		if !reflect.DeepEqual(result, testResult[i]) {
			t.Errorf(`Expected (%v)`, testResult[i])
		}
	}
}
