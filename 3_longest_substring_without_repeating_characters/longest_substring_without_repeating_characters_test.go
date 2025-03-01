package __longest_substring_without_repeating_characters

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Example 1: Basic case with repeating letters",
			input:    "abcabcbb",
			expected: 3,
		},
		{
			name:     "Example 2: All same characters",
			input:    "bbbbb",
			expected: 1,
		},
		{
			name:     "Example 3: Multiple non-consecutive repeating characters",
			input:    "pwwkew",
			expected: 3,
		},
		{
			name:     "Empty string",
			input:    "",
			expected: 0,
		},
		{
			name:     "Single character",
			input:    "a",
			expected: 1,
		},
		{
			name:     "Two different characters",
			input:    "ab",
			expected: 2,
		},
		{
			name:     "String with spaces",
			input:    "helloxworld",
			expected: 6,
		},
		{
			name:     "String with numbers",
			input:    "abc123xyz",
			expected: 9,
		},
		{
			name:     "String with special characters",
			input:    "!@#$%^&*()",
			expected: 10,
		},
		{
			name:     "Repeating pattern",
			input:    "abababab",
			expected: 2,
		},
		{
			name:     "Late repeating character",
			input:    "abcdefghija",
			expected: 10,
		},
		{
			name:     "All unique characters",
			input:    "abcdefghijklmnop",
			expected: 16,
		},
		{
			name:     "Mixed case letters",
			input:    "abcABC",
			expected: 6,
		},
		{
			name:     "Complex pattern with multiple types",
			input:    "1a2b3c4d5e1",
			expected: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lengthOfLongestSubstring(tt.input)
			if result != tt.expected {
				t.Errorf("lengthOfLongestSubstring(%q) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}
