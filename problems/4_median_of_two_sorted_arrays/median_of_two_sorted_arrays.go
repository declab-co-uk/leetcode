package __median_of_two_sorted_arrays

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	//Create lists
	var n, m []int

	//Set smaller list to n, larger to n
	if len(nums1) > len(nums2) {
		n = nums2
		m = nums1
	} else {
		n = nums1
		m = nums2
	}

	lenN, lenM := len(n), len(m)

	//Create binary search params
	left, right := 0, lenN

	for left <= right {
		splitN := (left + right) / 2
		splitM := ((lenM + lenN) / 2) - splitN

		nU := math.Pow(10, 6) + 1
		if splitN != lenN {
			nU = float64(n[splitN])
		}

		nL := -(math.Pow(10, 6)) - 1
		if splitN != 0 {
			nL = float64(n[splitN-1])
		}

		mU := math.Pow(10, 6) + 1
		if splitM != lenM {
			mU = float64(m[splitM])
		}

		mL := -(math.Pow(10, 6)) - 1
		if splitM != 0 {
			mL = float64(m[splitM-1])
		}

		if nU >= mL && mU >= nL {

			if (lenM+lenN)%2 == 0 {
				return (math.Max(mL, nL) + math.Min(mU, nU)) / 2
			} else {
				return math.Min(mU, nU)
			}
		}

		if nU <= mL {
			//increase nU decrease mL
			left = splitN + 1
		}
		if mU <= nL {
			//increase mU decrease nL
			right = splitN - 1

		}

	}

	return 0
}
