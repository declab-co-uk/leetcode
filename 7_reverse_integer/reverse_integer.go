package __reverse_integer

import (
	"math"
)

func reverse(x int) int {

	lenX := math.Floor(math.Log10(math.Abs(float64(x))))
	output := 0
	excess := 0

	for i := 0; i <= int(lenX); i++ {
		get := int(math.Pow(10, float64(i+1)))
		getDigit := ((x % get) - excess) / (get / 10)
		set := int(math.Pow(10, lenX-float64(i)))
		if int(lenX)-i < 1 {
			set = 1
		}
		output += (set * getDigit)
		excess += getDigit - excess
	}

	if output > math.MaxInt32 || output < math.MinInt32 {
		return 0
	}

	return output
}
