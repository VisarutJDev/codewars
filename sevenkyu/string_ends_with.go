package sevenkyu

func Solution(str, ending string) bool {
	// Your code here!
	lenOfStr := len(str)
	lenOfEnding := len(ending)
	if lenOfEnding == 0 {
		return true
	}

	if lenOfStr < lenOfEnding {
		return false
	}

	cut := str[lenOfStr-lenOfEnding:]
	return cut == ending
}
