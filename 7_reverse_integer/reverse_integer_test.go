package __reverse_integer

import (
	"testing"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		// Basic positive numbers
		{
			name:     "basic positive number 123",
			input:    123,
			expected: 321,
		},
		{
			name:     "basic positive number 456",
			input:    456,
			expected: 654,
		},

		// Numbers with trailing zeros
		{
			name:     "number with trailing zeros 120",
			input:    120,
			expected: 21,
		},
		{
			name:     "number with multiple trailing zeros 10000",
			input:    10000,
			expected: 1,
		},

		// Negative numbers
		{
			name:     "negative number -123",
			input:    -123,
			expected: -321,
		},
		{
			name:     "negative number with trailing zero -10",
			input:    -10,
			expected: -1,
		},

		// Single-digit numbers
		{
			name:     "single digit 0",
			input:    0,
			expected: 0,
		},
		{
			name:     "single digit 5",
			input:    5,
			expected: 5,
		},

		// Palindrome numbers
		{
			name:     "palindrome 121",
			input:    121,
			expected: 121,
		},
		{
			name:     "palindrome 12321",
			input:    12321,
			expected: 12321,
		},

		// Numbers with sequence of same digits
		{
			name:     "repeated digits 111",
			input:    111,
			expected: 111,
		},
		{
			name:     "repeated digits 9999",
			input:    9999,
			expected: 9999,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := reverse(tc.input)
			if result != tc.expected {
				t.Errorf("reverse(%d) = %d; expected %d", tc.input, result, tc.expected)
			}
		})
	}
}

// Additional benchmarking test
func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverse(123456789)
	}
}
