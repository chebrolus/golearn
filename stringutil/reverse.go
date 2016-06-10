package stringutil

func Reverse(input string) string {
	runeArray := []rune(input)
	for i, j := 0, len(runeArray)-1; i < len(runeArray)/2; i, j = i+1, j-1 {
		runeArray[i], runeArray[j] = runeArray[j], runeArray[i]
	}

	return string(runeArray)
}
