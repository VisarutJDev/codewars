package unittest

import (
	"reflect"
	"testing"

	"github.com/VisarutJDev/codewars/eightkyu"
)

func TestStringToNumber(t *testing.T) {
	testsample := []string{
		"1234",
		"605",
		"1405",
		"-7",
	}
	testResult := []int{
		1234,
		605,
		1405,
		-7,
	}
	for i := range testsample {
		result := eightkyu.StringToNumber(testsample[i])
		if !reflect.DeepEqual(result, testResult[i]) {
			t.Errorf(`Expected (%v)`, testResult[i])
		}
	}
}
