package __zigzag_conversion

import (
	"testing"
)

func TestConvert(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		numRows  int
		expected string
	}{
		// Examples from the problem statement
		{
			name:     "Example 1",
			input:    "PAYPALISHIRING",
			numRows:  3,
			expected: "PAHNAPLSIIGYIR",
		},
		{
			name:     "Example 2",
			input:    "PAYPALISHIRING",
			numRows:  4,
			expected: "PINALSIGYAHRPI",
		},
		{
			name:     "Example 3",
			input:    "A",
			numRows:  1,
			expected: "A",
		},

		// Edge cases
		{
			name:     "Empty string",
			input:    "",
			numRows:  3,
			expected: "",
		},
		{
			name:     "Single character with numRows = 2",
			input:    "A",
			numRows:  2,
			expected: "A",
		},
		{
			name:     "Two characters with numRows = 2",
			input:    "AB",
			numRows:  2,
			expected: "AB",
		},
		{
			name:     "String length equals numRows",
			input:    "ABC",
			numRows:  3,
			expected: "ABC",
		},
		{
			name:     "String length less than numRows",
			input:    "AB",
			numRows:  3,
			expected: "AB",
		},

		// Cases with special patterns
		{
			name:     "All characters in one row",
			input:    "ABCDEF",
			numRows:  1,
			expected: "ABCDEF",
		},
		{
			name:     "Two rows zigzag",
			input:    "ABCDEFGH",
			numRows:  2,
			expected: "ACEGBDFH",
		},
		{
			name:     "Three rows zigzag",
			input:    "ABCDEFGHIJK",
			numRows:  3,
			expected: "AEIBDFHJCGK",
		},

		// Boundary cases from constraints
		{
			name:     "Minimum numRows",
			input:    "ABCDEFG",
			numRows:  1,
			expected: "ABCDEFG",
		},

		{
			name:     "numRows equals string length",
			input:    "HELLO",
			numRows:  5,
			expected: "HELLO",
		},
		{
			name:     "numRows greater than string length",
			input:    "HELLO",
			numRows:  10,
			expected: "HELLO",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := convert(tc.input, tc.numRows)
			if result != tc.expected {
				t.Errorf("For input %q with %d rows: expected %q, got %q",
					tc.input, tc.numRows, tc.expected, result)
			}
		})
	}
}
