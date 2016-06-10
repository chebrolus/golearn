package stringutil

func isPalindrome(input string) bool {
	length := len(input)
	if length == 0 {
		return false
	}
	for i := 0; i < length/2; i++ {
		if input[i] != input[length-i-1] {
			return false
		}
	}
	return true
}
