package solution

import (
	"slices"
	"testing"
)

func TestRightSideView(t *testing.T) {
	// Explicit links keep sparse-tree fixtures unambiguous.
	node := func(value int, left, right *TreeNode) *TreeNode {
		return &TreeNode{Val: value, Left: left, Right: right}
	}

	var chain *TreeNode
	chainWant := make([]int, 100)
	for value := 99; value >= 0; value-- {
		chain = node(value, chain, nil)
		chainWant[value] = value
	}

	tests := []struct {
		name string
		root *TreeNode
		want []int
	}{
		{
			name: "example sparse tree",
			root: node(1, node(2, nil, node(5, nil, nil)), node(3, nil, node(4, nil, nil))),
			want: []int{1, 3, 4},
		},
		{
			name: "example left subtree extends deeper",
			root: node(1, node(2, node(4, node(5, nil, nil), nil), nil), node(3, nil, nil)),
			want: []int{1, 3, 4, 5},
		},
		{name: "example right child only", root: node(1, nil, node(3, nil, nil)), want: []int{1, 3}},
		{name: "empty tree", root: nil, want: []int{}},
		{name: "single node", root: node(42, nil, nil), want: []int{42}},
		{
			name: "complete tree",
			root: node(1, node(2, node(4, nil, nil), node(5, nil, nil)), node(3, node(6, nil, nil), node(7, nil, nil))),
			want: []int{1, 3, 7},
		},
		{
			name: "right subtree left child beats left subtree right child",
			root: node(1, node(2, nil, node(5, nil, nil)), node(3, node(6, nil, nil), nil)),
			want: []int{1, 3, 6},
		},
		{
			name: "zigzag chain",
			root: node(1, nil, node(2, node(3, nil, node(4, nil, nil)), nil)),
			want: []int{1, 2, 3, 4},
		},
		{
			name: "repeated values at different depths",
			root: node(7, node(7, node(7, nil, nil), nil), node(7, nil, nil)),
			want: []int{7, 7, 7},
		},
		{name: "negative zero and boundary values", root: node(0, node(100, nil, nil), node(-100, nil, nil)), want: []int{0, -100}},
		{name: "maximum length left chain", root: chain, want: chainWant},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rightSideView(tt.root); !slices.Equal(got, tt.want) {
				t.Fatalf("rightSideView() = %v; want %v", got, tt.want)
			}
		})
	}
}
