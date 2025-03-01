package __add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	//Create initial carry value
	carry := 0
	//Start recursion
	return addNodes(l1, l2, carry)

}

func addNodes(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	//Create next node
	outputNode := new(ListNode)
	//Create empty list pointers to avoid nil dereferencing in recursion loop
	var nextL1, nextL2 *ListNode

	sum := 0
	if l1 != nil {
		//Add L1 value and prepare the next list value
		sum += l1.Val
		nextL1 = l1.Next
	}
	if l2 != nil {
		//Add L2 value and prepare the next list value
		sum += l2.Val
		nextL2 = l2.Next
	}
	//Add the carry
	sum += carry

	//if the total is 10 or above, carry the 10 and set val to the remainder
	if sum > 9 {
		outputNode.Val = sum % 10
		carry = 1
	} else {
		outputNode.Val = sum
		carry = 0
	}

	// if the lists are not exhausted and carry is not 0, continue recursion
	if nextL1 != nil || nextL2 != nil || carry != 0 {
		outputNode.Next = addNodes(nextL1, nextL2, carry)
	}

	return outputNode

}
