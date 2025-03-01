package __longest_palindromic_substring

func longestPalindrome(s string) string {

	currentLongest := ""

	for i := range s {
		outputString := getPalindrome(&s, int(i), int(i))
		if len(outputString) > len(currentLongest) {
			currentLongest = outputString
		}
	}

	return currentLongest
}

func getPalindrome(s *string, left, right int) string {
	if left == right {
		for right != len(*s)-1 {
			if (*s)[right] == (*s)[right+1] {
				right += 1
			} else {
				break
			}
		}
	}
	if (*s)[left] != (*s)[right] {
		return (*s)[left+1 : right]
	}
	if left == 0 || right == len(*s)-1 {
		return (*s)[left : right+1]
	}

	return getPalindrome(s, left-1, right+1)
}
