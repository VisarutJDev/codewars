package unittest

import (
	"reflect"
	"testing"

	"github.com/VisarutJDev/codewars/eightkyu"
)

func TestNumberToString(t *testing.T) {
	testsample := []int{
		67,
		79585,
		79585,
		3,
		-1,
	}
	testResult := []string{
		"67",
		"79585",
		"79585",
		"3",
		"-1",
	}
	for i := range testsample {
		result := eightkyu.NumberToString(testsample[i])
		if !reflect.DeepEqual(result, testResult[i]) {
			t.Errorf(`Expected (%v)`, testResult[i])
		}
	}
}
