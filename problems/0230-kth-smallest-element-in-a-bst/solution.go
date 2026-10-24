package solution

// kthSmallest returns the value at the 1-indexed rank k without modifying the BST.
// The problem guarantees a nonempty tree and 1 <= k <= the number of nodes.
func kthSmallest(root *TreeNode, k int) int {
	stack := make([]*TreeNode, 0)
	current := root

	for current != nil || len(stack) > 0 {
		// Defer each ancestor until its smaller, left-subtree values are visited.
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return current.Val
		}
		current = current.Right
	}

	panic("kthSmallest: k is outside the tree's ranks")
}
