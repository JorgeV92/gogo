package solution

// rightSideView returns the rightmost value at each depth, from top to bottom.
// It visits the right subtree first and does not modify the tree.
func rightSideView(root *TreeNode) []int {
	view := make([]int, 0)

	var visit func(node *TreeNode, depth int)
	visit = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		// The first node reached at this depth is the rightmost one.
		if depth == len(view) {
			view = append(view, node.Val)
		}
		visit(node.Right, depth+1)
		visit(node.Left, depth+1)
	}

	visit(root, 0)
	return view
}
