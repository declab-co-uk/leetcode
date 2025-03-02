package __zigzag_conversion

func convert(s string, numRows int) string {
	cycleLength := (numRows * 2) - 2
	charMap := make(map[int][]byte)
	if numRows == 1 {
		return s
	}

	for i := 0; i < len(s); i++ {
		positionInCycle := i % cycleLength
		if positionInCycle < numRows {
			charMap[positionInCycle] = append(charMap[positionInCycle], s[i])
		} else {
			n := numRows - 1
			m := positionInCycle - n
			row := n - m
			charMap[row] = append(charMap[row], s[i])
		}
	}

	outputString := make([]byte, 0)
	for i := 0; i <= numRows-1; i++ {
		outputString = append(outputString, charMap[i]...)
	}

	return string(outputString)

}
