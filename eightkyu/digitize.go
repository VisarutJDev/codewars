package eightkyu

// Convert number to reversed array of digits
// Given a random non-negative number, you have to return the digits of this number within an array in reverse order.

// Example(Input => Output):
// 35231 => [1,3,2,5,3]
// 0 => [0]
func Digitize(n int) []int {
	// your code here
	splitNumber := []int{}
	if n/10 == 0 {
		return []int{n}
	}
	for n > 0 {
		splitNumber = append(splitNumber, n%10)
		n = n / 10
	}

	return splitNumber
}
