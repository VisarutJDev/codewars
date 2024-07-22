package unittest

import (
	"reflect"
	"testing"

	"github.com/VisarutJDev/codewars/sevenkyu"
)

func TestReverseWords(t *testing.T) {
	testsample := []string{
		"The quick brown fox jumps over the lazy dog.",
		"apple",
		"a b c d",
		"double  spaced  words",
		"stressed desserts",
		"hello hello",
	}
	testResult := []string{
		"ehT kciuq nworb xof spmuj revo eht yzal .god",
		"elppa",
		"a b c d",
		"elbuod  decaps  sdrow",
		"desserts stressed",
		"olleh olleh",
	}
	for i := range testsample {
		result := sevenkyu.ReverseWords(testsample[i])
		if !reflect.DeepEqual(result, testResult[i]) {
			t.Errorf(`Expected (%v)`, testResult[i])
		}
	}

}
