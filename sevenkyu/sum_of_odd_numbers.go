package sevenkyu

// 			   1
// 			3     5
// 	 	 7     9    11
// 	  13    15    17    19
// 21    23    25    27    29
// ...
// n = n
// 1 = n+0 ; n =1
// 8 = n+1 + n+1+2 ; n=2
// 27 = n+1+3 n+1+5 n+1+7 ; n = 3
// 64 = n+1+8   ; n = 4
// RowSumOddNumbers Sum of odd numbers
func RowSumOddNumbers(n int) int {
	// your code here
	return n * n * n
}
