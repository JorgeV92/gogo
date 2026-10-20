package solution

import "sort"

// getStrongest returns the k strongest values, sorting arr in place.
// It assumes LeetCode's constraints: len(arr) >= 1 and 1 <= k <= len(arr).
func getStrongest(arr []int, k int) []int {
	sort.Ints(arr)
	median := arr[(len(arr)-1)/2]
	left, right := 0, len(arr)-1
	strongest := make([]int, 0, k)

	for len(strongest) < k {
		// The strongest remaining value is at an endpoint. On equal
		// distances, take the right endpoint because its value is larger.
		if median-arr[left] > arr[right]-median {
			strongest = append(strongest, arr[left])
			left++
		} else {
			strongest = append(strongest, arr[right])
			right--
		}
	}

	return strongest
}
