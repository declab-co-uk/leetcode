package __add_two_numbers

import "testing"

// Helper function to compare two linked lists
func compareLinkedLists(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	// Check if both lists have reached their end
	return l1 == nil && l2 == nil
}

// Helper function to convert linked list to slice for better error messages
func linkedListToSlice(l *ListNode) []int {
	var result []int
	for l != nil {
		result = append(result, l.Val)
		l = l.Next
	}
	return result
}

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{
			name: "Example 1: 342 + 465 = 807",
			l1:   &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
			l2:   &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
			want: &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8}}},
		},
		{
			name: "Example 2: 0 + 0 = 0",
			l1:   &ListNode{Val: 0},
			l2:   &ListNode{Val: 0},
			want: &ListNode{Val: 0},
		},
		{
			name: "Example 3: 9999999 + 9999",
			l1:   &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}}}},
			l2:   &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}},
			want: &ListNode{Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}}}}}}},
		},
		{
			name: "Different length lists: 99 + 1",
			l1:   &ListNode{Val: 9, Next: &ListNode{Val: 9}},
			l2:   &ListNode{Val: 1},
			want: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}},
		},
		{
			name: "Simple carry: 5 + 5 = 10",
			l1:   &ListNode{Val: 5},
			l2:   &ListNode{Val: 5},
			want: &ListNode{Val: 0, Next: &ListNode{Val: 1}},
		},
		{
			name: "Zero Value: 81 + 0 = 81",
			l1:   &ListNode{Val: 1, Next: &ListNode{Val: 8}},
			l2:   &ListNode{Val: 0},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 8}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addTwoNumbers(tt.l1, tt.l2)
			if !compareLinkedLists(got, tt.want) {
				t.Errorf("AddTwoNumbers() = %v, want %v", linkedListToSlice(got), linkedListToSlice(tt.want))
			}
		})
	}
}
