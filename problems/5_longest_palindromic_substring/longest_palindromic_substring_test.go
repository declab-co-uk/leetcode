package __longest_palindromic_substring

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		// Basic test cases from the problem statement
		{
			name:     "Example 1 - Multiple palindromes of same length",
			input:    "babad",
			expected: "bab", // or "aba"
		},
		{
			name:     "Example 2 - Even length palindrome",
			input:    "cbbd",
			expected: "bb",
		},

		// Additional test cases
		{
			name:     "Single character",
			input:    "a",
			expected: "a",
		},
		{
			name:     "Two identical characters",
			input:    "aa",
			expected: "aa",
		},
		{
			name:     "No palindrome longer than 1",
			input:    "abc",
			expected: "a", // or "b" or "c"
		},
		{
			name:     "Entire string is a palindrome",
			input:    "racecar",
			expected: "racecar",
		},
		{
			name:     "Longest palindrome at beginning",
			input:    "abbcdefg",
			expected: "bb",
		},
		{
			name:     "Longest palindrome at end",
			input:    "defgabb",
			expected: "bb",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Long string with multiple palindromes",
			input:    "civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth",
			expected: "ranynar", // one of the longest palindromes in this string
		},
		{
			name:     "Multiple digits and letters",
			input:    "1a2b3ccc3b2a1",
			expected: "1a2b3ccc3b2a1",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := longestPalindrome(tc.input)

			// For cases where multiple answers are valid, we need a more flexible check
			if tc.name == "Example 1 - Multiple palindromes of same length" {
				if result != "bab" && result != "aba" {
					t.Errorf("expected either 'bab' or 'aba', got %s", result)
				}
			} else if tc.name == "No palindrome longer than 1" {
				if len(result) != 1 || !contains(tc.input, result) {
					t.Errorf("expected a single character from input, got %s", result)
				}
			} else if len(result) != len(tc.expected) || !isPalindrome(result) {
				t.Errorf("expected palindrome of length %d, got %s with length %d",
					len(tc.expected), result, len(result))
			}
		})
	}
}

// Helper function to check if a string is a palindrome
func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

// Helper function to check if string contains substring
func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// Benchmark test to evaluate performance
func BenchmarkLongestPalindrome(b *testing.B) {
	input := "civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendure"
	for i := 0; i < b.N; i++ {
		longestPalindrome(input)
	}
}
