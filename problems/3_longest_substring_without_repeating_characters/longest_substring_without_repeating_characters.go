package __longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {

	if len(s) <= 1 {
		return len(s)
	}

	windowStart := 0
	currentLongest := 0
	indexMap := make(map[rune]int)

	for windowEnd, character := range s {
		if index, exists := indexMap[character]; exists && index >= windowStart {
			windowStart = index + 1
		}
		indexMap[character] = windowEnd
		if windowEnd-windowStart+1 > currentLongest {
			currentLongest = windowEnd - windowStart + 1
		}

	}

	return currentLongest

}
