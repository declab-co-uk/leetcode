package __median_of_two_sorted_arrays

import (
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  float64
	}{
		// Basic cases
		{
			name:  "Example 1 from problem statement",
			nums1: []int{1, 3},
			nums2: []int{2},
			want:  2.0,
		},
		{
			name:  "Example 2 from problem statement",
			nums1: []int{1, 2},
			nums2: []int{3, 4},
			want:  2.5,
		},

		// Edge cases with empty arrays
		{
			name:  "First array empty",
			nums1: []int{},
			nums2: []int{1},
			want:  1.0,
		},
		{
			name:  "Second array empty",
			nums1: []int{1},
			nums2: []int{},
			want:  1.0,
		},

		// Single element arrays
		{
			name:  "Both single element",
			nums1: []int{1},
			nums2: []int{2},
			want:  1.5,
		},

		// Arrays with different lengths
		{
			name:  "First array longer",
			nums1: []int{1, 2, 3, 4, 5},
			nums2: []int{6, 7, 8},
			want:  4.5,
		},
		{
			name:  "Second array longer",
			nums1: []int{1, 2, 3},
			nums2: []int{4, 5, 6, 7, 8},
			want:  4.5,
		},

		// Arrays with negative numbers
		{
			name:  "Negative numbers",
			nums1: []int{-5, -3, -1},
			nums2: []int{-2, 0, 2},
			want:  -1.5,
		},

		// Arrays with duplicate numbers
		{
			name:  "Duplicates across arrays",
			nums1: []int{1, 2, 2},
			nums2: []int{2, 3, 4},
			want:  2.0,
		},

		// Large numbers (close to constraints)
		{
			name:  "Large positive numbers",
			nums1: []int{100000, 100001},
			nums2: []int{100002, 100003},
			want:  100001.5,
		},
		{
			name:  "Large negative numbers",
			nums1: []int{-100003, -100002},
			nums2: []int{-100001, -100000},
			want:  -100001.5,
		},

		// Boundary cases for array lengths
		{
			name:  "Large first array",
			nums1: makeArray(1000, 1), // 1000 elements starting from 1
			nums2: []int{1001},
			want:  501,
		},
		{
			name:  "Large second array",
			nums1: []int{0},
			nums2: makeArray(1000, 1), // 1000 elements starting from 1
			want:  500.0,
		},

		// Special patterns
		{
			name:  "Alternating pattern",
			nums1: []int{1, 3, 5, 7},
			nums2: []int{2, 4, 6, 8},
			want:  4.5,
		},
		{
			name:  "One array entirely greater",
			nums1: []int{1, 2, 3},
			nums2: []int{4, 5, 6},
			want:  3.5,
		},

		// Edge cases with identical numbers
		{
			name:  "All same numbers",
			nums1: []int{2, 2, 2},
			nums2: []int{2, 2},
			want:  2.0,
		},

		// Mixed positive and negative
		{
			name:  "Mixed positive negative with zero",
			nums1: []int{-5, -2, 0, 3, 7},
			nums2: []int{-3, -1, 2, 4},
			want:  0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMedianSortedArrays(tt.nums1, tt.nums2)
			if got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Helper function to create large arrays for testing
func makeArray(size int, start int) []int {
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = start + i
	}
	return result
}
