package solution

// maxSumThreeSubarrays returns the starting indices of three non-overlapping
// length-k subarrays with maximum total sum, breaking ties lexicographically.
func maxSumThreeSubarrays(nums []int, k int) []int {
	windowCount := len(nums) - k + 1
	windowSums := make([]int, windowCount)
	for i := 0; i < k; i++ {
		windowSums[0] += nums[i]
	}
	for i := 1; i < windowCount; i++ {
		windowSums[i] = windowSums[i-1] - nums[i-1] + nums[i+k-1]
	}

	left := make([]int, windowCount)
	for i := 1; i < windowCount; i++ {
		left[i] = left[i-1]
		if windowSums[i] > windowSums[left[i]] {
			left[i] = i
		}
	}

	right := make([]int, windowCount)
	right[windowCount-1] = windowCount - 1
	for i := windowCount - 2; i >= 0; i-- {
		right[i] = right[i+1]
		if windowSums[i] >= windowSums[right[i]] {
			right[i] = i
		}
	}

	answer := []int{0, k, 2 * k}
	maxTotal := -1
	for mid := k; mid <= len(nums)-2*k; mid++ {
		first, third := left[mid-k], right[mid+k]
		total := windowSums[first] + windowSums[mid] + windowSums[third]
		if total > maxTotal {
			maxTotal = total
			answer = []int{first, mid, third}
		}
	}
	return answer
}
