package unittest

import (
	"reflect"
	"testing"

	"github.com/VisarutJDev/codewars/eightkyu"
)

func TestLitres(t *testing.T) {
	testsample := []float64{
		2,
		1.4,
		12.3,
		0.82,
		11.8,
		1787,
		0,
	}
	testResult := []int{
		1,
		0,
		6,
		0,
		5,
		893,
		0,
	}
	for i := range testsample {
		result := eightkyu.Litres(testsample[i])
		if !reflect.DeepEqual(result, testResult[i]) {
			t.Errorf(`Expected (%v)`, testResult[i])
		}
	}

}
