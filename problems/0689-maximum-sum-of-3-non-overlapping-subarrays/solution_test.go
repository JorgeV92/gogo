package solution

import (
	"reflect"
	"testing"
)

func TestMaxSumThreeSubarrays(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{name: "example one", nums: []int{1, 2, 1, 2, 6, 7, 5, 1}, k: 2, want: []int{0, 3, 5}},
		{name: "example two", nums: []int{1, 2, 1, 2, 1, 2, 1, 2, 1}, k: 2, want: []int{0, 2, 4}},
		{name: "minimum input length", nums: []int{1, 2, 3, 4, 5, 6}, k: 2, want: []int{0, 2, 4}},
		{name: "all equal", nums: []int{7, 7, 7, 7, 7, 7, 7, 7, 7}, k: 1, want: []int{0, 1, 2}},
		{name: "separated maxima", nums: []int{1, 1, 10, 1, 1, 20, 1, 1, 15}, k: 1, want: []int{2, 5, 8}},
		{name: "equal maxima choose earliest", nums: []int{1, 1, 5, 1, 1, 5, 1, 1, 5}, k: 1, want: []int{2, 5, 8}},
		{name: "larger windows", nums: []int{4, 5, 10, 1, 2, 8, 9, 3, 7, 6, 1, 1}, k: 3, want: []int{0, 4, 7}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := append([]int(nil), tt.nums...)
			got := maxSumThreeSubarrays(input, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("maxSumThreeSubarrays(%v, %d) = %v; want %v", tt.nums, tt.k, got, tt.want)
			}
			if !reflect.DeepEqual(input, tt.nums) {
				t.Fatalf("maxSumThreeSubarrays modified its input: got %v; want %v", input, tt.nums)
			}
		})
	}
}

func TestMaxSumThreeSubarraysMaximumLength(t *testing.T) {
	const n = 20000
	const k = 6000
	nums := make([]int, n)
	for i := range nums {
		nums[i] = 1
	}

	want := []int{0, k, 2 * k}
	if got := maxSumThreeSubarrays(nums, k); !reflect.DeepEqual(got, want) {
		t.Fatalf("maxSumThreeSubarrays(maximum-length input, %d) = %v; want %v", k, got, want)
	}
}
