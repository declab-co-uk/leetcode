package __string_to_integer_atoi

import "testing"

func TestAtoi(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "9223372036854775808",
			input:    "9223372036854775808",
			expected: 2147483647,
		},
		{
			name:     "   -042",
			input:    "   -042",
			expected: -42,
		},
		{
			name:     "21474836460",
			input:    "21474836460",
			expected: 2147483647,
		},
		{
			name:     "-91283472332",
			input:    "-91283472332",
			expected: -2147483648,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := myAtoi(tc.input)
			if result != tc.expected {
				t.Errorf("Atoi(%s) = %d; expected %d", tc.input, result, tc.expected)
			}

		})
	}
}
