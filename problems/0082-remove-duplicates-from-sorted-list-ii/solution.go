package solution

// deleteDuplicates removes every occurrence of values that appear more than once.
// It assumes a sorted list and modifies its links, reusing the surviving nodes.
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := ListNode{Next: head}
	prev := &dummy
	current := head

	for current != nil {
		if current.Next != nil && current.Val == current.Next.Val {
			duplicate := current.Val
			// Skip the entire run, including its first node.
			for current != nil && current.Val == duplicate {
				current = current.Next
			}
			prev.Next = current
		} else {
			// This value occurs once; retain it and advance the output tail.
			prev = current
			current = current.Next
		}
	}

	return dummy.Next
}
