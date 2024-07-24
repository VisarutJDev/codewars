package eightkyu

import "strconv"

func StringToNumber(str string) int {
	// your code here
	r, _ := strconv.Atoi(str)
	return r
}
