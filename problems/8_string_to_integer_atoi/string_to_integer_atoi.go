package __string_to_integer_atoi

import (
	"math"
	"strings"
)

func myAtoi(s string) int {

	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return 0
	}

	var isNegative bool

	if s[0] == '-' || s[0] == '+' {
		if s[0] == '-' {
			isNegative = true
		}
		s = s[1:]
	}

	output := 0

	for index, char := range s {
		if char < '0' || char > '9' {
			break
		}

		output = output*10 + int(s[index]-'0')

		if output > math.MaxInt32 && !isNegative {
			return math.MaxInt32
		}
		if -output < math.MinInt32 && isNegative {
			return math.MinInt32
		}
	}

	if isNegative {
		output = -output
	}

	return output
}
