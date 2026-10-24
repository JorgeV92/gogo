package solution

import "testing"

func TestKthSmallest(t *testing.T) {
	node := func(value int, left, right *TreeNode) *TreeNode {
		return &TreeNode{Val: value, Left: left, Right: right}
	}

	balanced := node(40, node(10, node(0, nil, nil), node(25, nil, nil)), node(90, node(60, nil, nil), node(10000, nil, nil)))
	sparse := node(8, node(3, nil, node(6, node(4, nil, nil), node(7, nil, nil))), node(10, nil, node(14, node(13, nil, nil), nil)))
	var leftChain, rightChain *TreeNode
	for value := 0; value < 10000; value++ {
		leftChain = node(value, leftChain, nil)
		rightChain = node(9999-value, nil, rightChain)
	}

	tests := []struct {
		name string
		root *TreeNode
		k    int
		want int
	}{
		{
			name: "first example",
			root: node(3, node(1, nil, node(2, nil, nil)), node(4, nil, nil)),
			k:    1,
			want: 1,
		},
		{
			name: "second example",
			root: node(5, node(3, node(2, node(1, nil, nil), nil), node(4, nil, nil)), node(6, nil, nil)),
			k:    3,
			want: 3,
		},
		{name: "single zero", root: node(0, nil, nil), k: 1, want: 0},
		{name: "minimum", root: balanced, k: 1, want: 0},
		{name: "left subtree interior", root: balanced, k: 3, want: 25},
		{name: "root rank", root: balanced, k: 4, want: 40},
		{name: "right subtree minimum", root: balanced, k: 5, want: 60},
		{name: "maximum", root: balanced, k: 7, want: 10000},
		{name: "left descent after right turn", root: sparse, k: 2, want: 4},
		{name: "ancestor after completed subtree", root: sparse, k: 5, want: 8},
		{name: "sparse right subtree", root: sparse, k: 7, want: 13},
		{name: "maximum length left chain first", root: leftChain, k: 1, want: 0},
		{name: "maximum length left chain last", root: leftChain, k: 10000, want: 9999},
		{name: "maximum length right chain first", root: rightChain, k: 1, want: 0},
		{name: "maximum length right chain middle", root: rightChain, k: 5000, want: 4999},
		{name: "maximum length right chain last", root: rightChain, k: 10000, want: 9999},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest(tt.root, tt.k); got != tt.want {
				t.Fatalf("kthSmallest(root, %d) = %d; want %d", tt.k, got, tt.want)
			}
		})
	}
}
