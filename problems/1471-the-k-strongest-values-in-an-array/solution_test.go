package solution

import (
	"reflect"
	"sort"
	"testing"
)

func TestGetStrongest(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		k    int
		want []int
	}{
		{name: "odd length", arr: []int{1, 2, 3, 4, 5}, k: 2, want: []int{5, 1}},
		{name: "duplicates", arr: []int{1, 1, 3, 5, 5}, k: 2, want: []int{5, 5}},
		{name: "unsorted even length", arr: []int{6, 7, 11, 7, 6, 8}, k: 5, want: []int{11, 8, 6, 6, 7}},
		{name: "tie at selection boundary", arr: []int{1, 3, 5}, k: 1, want: []int{5}},
		{name: "lower median", arr: []int{1, 2, 3, 4}, k: 1, want: []int{4}},
		{name: "negative values", arr: []int{-9, -5, -3, -2, -1}, k: 2, want: []int{-9, -1}},
		{name: "single value", arr: []int{42}, k: 1, want: []int{42}},
		{name: "all equal", arr: []int{7, 7, 7, 7}, k: 3, want: []int{7, 7, 7}},
		{name: "take all", arr: []int{4, -2, 9, 1}, k: 4, want: []int{9, 4, -2, 1}},
		{name: "constraint extremes", arr: []int{-100000, 0, 100000}, k: 1, want: []int{100000}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// The implementation sorts its input; preserve the original fixture.
			input := append([]int(nil), tt.arr...)
			got := getStrongest(input, tt.k)
			want := append([]int(nil), tt.want...)
			// Any output order is valid, but duplicate counts must match.
			sort.Ints(got)
			sort.Ints(want)
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("getStrongest(%v, %d) = %v; want %v (order ignored)", tt.arr, tt.k, got, want)
			}
		})
	}
}
