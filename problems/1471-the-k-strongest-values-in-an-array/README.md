# 1471. The k Strongest Values in an Array

| Field | Value |
| --- | --- |
| Problem | [LeetCode 1471](https://leetcode.com/problems/the-k-strongest-values-in-an-array/) |
| Source | [Google question list](https://github.com/snehasishroy/leetcode-companywise-interview-questions/blob/master/google/all.csv) |
| Difficulty | Medium |
| Acceptance / frequency | 62.8% / 25.0% (supplied source snapshot) |
| Patterns | Sorting, two pointers |
| Status | Solved (example) |

## Problem in my own words

Choose `k` array entries ranked by distance from the median. Greater distance wins; equal distances favor the larger value. The median is the sorted element at index `(n-1)/2`, including for even lengths. Return the selected values in any order, preserving duplicates. See the [full statement and constraints](https://leetcode.com/problems/the-k-strongest-values-in-an-array/description/).

## Approach

Sorting the values locates the median. A second sort by strength would work, but the strongest remaining value must already be at one of the sorted array's two ends.

1. Sort `arr` in ascending order and save its lower median.
2. Place `left` and `right` at the endpoints.
3. Compare `median - arr[left]` with `arr[right] - median`.
4. Select the left value only when its distance is strictly greater; otherwise select the right value to resolve ties in favor of the larger number.
5. Move the selected pointer inward and repeat until `k` values are collected.

The implementation in [solution.go](solution.go) sorts the input in place. To preserve the caller's input, copy the slice before calling it; that adds `O(n)` storage.

## Walkthrough

For `arr = [8, 2, 5, 1, 10]` and `k = 3`, sorting gives `[1, 2, 5, 8, 10]` and median `5`.

| Step | Left value / distance | Right value / distance | Pick | Result |
| --- | --- | --- | --- | --- |
| 1 | 1 / 4 | 10 / 5 | 10 | `[10]` |
| 2 | 1 / 4 | 8 / 3 | 1 | `[10, 1]` |
| 3 | 2 / 3 | 8 / 3 | 8 (larger value wins tie) | `[10, 1, 8]` |

## Why it works

The unselected entries always form a contiguous interval of the sorted array. Distance from the fixed median is greatest at an endpoint of that interval, so comparing endpoints finds the next strongest value. If their distances tie, the right endpoint is at least as large and wins the tie. Removing the chosen endpoint preserves the interval. Repeating this choice selects exactly the strongest `k` entries.

The code uses signed differences rather than an absolute-value helper. While the interval straddles the median, both differences are nonnegative and are the actual distances. If all remaining values are above the median, the left expression becomes negative and the right endpoint correctly wins; below the median, the reverse holds. The median is saved once and never recomputed.

## Complexity

Let `n = len(arr)`.

- **Time: `O(n log n + k)`, or `O(n log n)` since `k <= n`.** Sorting dominates; each selection takes constant time.
- **Auxiliary space excluding the result: `O(log n)`** for the Go sorting implementation's stack. The pointer scan itself uses `O(1)` auxiliary space.
- **Output space: `O(k)`** for the returned slice. Including the sorting stack, total extra space is `O(log n + k)`.

These are asymptotic bounds, not measured LeetCode execution times.

## Edge cases

[Tests](solution_test.go) cover odd and even lengths, the lower-median rule, a tie at the selection boundary, duplicate values, negative values, a single element, all equal values, selecting the entire array, and extreme allowed values. They compare results without requiring a particular order, while checking duplicate counts.

The function assumes valid LeetCode inputs: a nonempty array and `1 <= k <= n`.

## Run

From the repository root:

```sh
go test -v ./problems/1471-the-k-strongest-values-in-an-array
```

## Interview follow-ups

- To preserve the input, sort a copy and account for its additional `O(n)` space.
- Quickselect can find the median and partition by strength in expected `O(n)` time, but requires more involved partition logic and has `O(n²)` worst-case time with naive pivot selection.
