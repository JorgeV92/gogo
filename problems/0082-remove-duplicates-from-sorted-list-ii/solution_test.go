package solution

import "testing"

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want []int
	}{
		{name: "middle duplicate runs", vals: []int{1, 2, 3, 3, 4, 4, 5}, want: []int{1, 2, 5}},
		{name: "duplicates at head", vals: []int{1, 1, 1, 2, 3}, want: []int{2, 3}},
		{name: "empty list", vals: nil, want: nil},
		{name: "single node", vals: []int{7}, want: []int{7}},
		{name: "all distinct", vals: []int{1, 2, 3}, want: []int{1, 2, 3}},
		{name: "equal pair", vals: []int{2, 2}, want: nil},
		{name: "all equal", vals: []int{5, 5, 5, 5}, want: nil},
		{name: "duplicates at tail", vals: []int{1, 2, 3, 3}, want: []int{1, 2}},
		{name: "every value duplicated", vals: []int{1, 1, 2, 2, 3, 3}, want: nil},
		{name: "only middle survives", vals: []int{1, 1, 2, 3, 3}, want: []int{2}},
		{name: "negative zero and boundary values", vals: []int{-100, -100, -1, 0, 0, 100}, want: []int{-1, 100}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dummy := ListNode{}
			tail := &dummy
			original := make(map[*ListNode]bool)
			for _, value := range tt.vals {
				tail.Next = &ListNode{Val: value}
				tail = tail.Next
				original[tail] = true
			}

			got := deleteDuplicates(dummy.Next)
			// Walk only the expected length, so a cycle fails instead of hanging.
			for i, want := range tt.want {
				if got == nil {
					t.Fatalf("node %d missing; want value %d", i, want)
				}
				if got.Val != want {
					t.Fatalf("node %d = %d; want %d", i, got.Val, want)
				}
				if !original[got] {
					t.Fatalf("node %d was allocated instead of reusing an input node", i)
				}
				got = got.Next
			}
			if got != nil {
				t.Fatalf("expected list to end after %d nodes; found extra node with value %d", len(tt.want), got.Val)
			}
		})
	}
}
