package __two_sum

func TwoSum(nums []int, target int) []int {
	//Slice to store the indices of the output numbers
	var outputIndices []int
	//Map to store the values where map[required_number]required_number_index
	var valueMap = make(map[int]int)
	//Iterate through num input
	for index, value := range nums {
		//Calculate the number to be added to the current number to equal the target
		requiredNumber := target - value
		//Check if the required number exists in the map
		result, exists := valueMap[requiredNumber]
		if exists {
			//If the required number is in the map, return both current and map index
			outputIndices = append(outputIndices, index, result)
			break
		}
		//If the value is not in the index we will add it here
		valueMap[value] = index
	}
	return outputIndices
}
