package sevenkyu

import (
	"strings"
)

func ReverseWords(str string) string {
	arr := strings.Split(str, " ")
	var arrResult []string
	for i := range arr {
		word := ""
		for j := len(arr[i]) - 1; j >= 0; j-- {
			word = word + string(arr[i][j])
		}
		arrResult = append(arrResult, word)
	}
	return strings.Join(arrResult[:], " ") // reverse those words
}
