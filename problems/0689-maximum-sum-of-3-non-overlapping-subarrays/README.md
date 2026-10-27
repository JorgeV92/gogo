# 689. Maximum Sum of 3 Non-Overlapping Subarrays

| Field | Value |
| --- | --- |
| Problem | [LeetCode 689](https://leetcode.com/problems/maximum-sum-of-3-non-overlapping-subarrays/) |
| Source | [Google question list](https://github.com/snehasishroy/leetcode-companywise-interview-questions/blob/master/google/all.csv) |
| Difficulty | Hard |
| Acceptance / frequency | 59.9% / 25.0% (supplied source snapshot) |
| Patterns | Dynamic programming, prefix sums, sliding window |
| Status | Solved |

## Problem in my own words

Choose three non-overlapping subarrays, each of length `k`, in their original left-to-right order. Maximize the sum of all their elements and return their three starting indices. If several choices have the same maximum sum, return the lexicographically smallest index list. The input has between 1 and 20,000 positive integers, and `k` is at most one third of its length. See the [full statement and constraints](https://leetcode.com/problems/maximum-sum-of-3-non-overlapping-subarrays/description/).

## Approach

Checking every triple of starting indices would take `O(n³)` time. Instead, treat the middle subarray as the pivot and precompute the best possible subarray wholly to its left and wholly to its right.

1. Compute every length-`k` window sum with a sliding window.
2. Build `left[i]`, the starting index of the maximum-sum window among starts `0` through `i`. On equal sums, retain the earlier index.
3. Build `right[i]`, the starting index of the maximum-sum window among starts `i` through the last possible start. On equal sums, retain the earlier index.
4. For each valid middle start `mid`, combine `left[mid-k]`, `mid`, and `right[mid+k]`. Update the answer only for a strictly larger total so equal totals retain the lexicographically smallest triple.

The implementation in [solution.go](solution.go) does not modify the input.

## Walkthrough

For `nums = [1,2,1,2,6,7,5,1]` and `k = 2`, the window sums by starting index are:

| Start | 0 | 1 | 2 | 3 | 4 | 5 | 6 |
| --- | --- | --- | --- | --- | --- | --- | --- |
| Sum | 3 | 3 | 3 | 8 | 13 | 12 | 6 |

For middle start `3`, the best left window starts at `0` and the best right window starts at `5`, giving total `3 + 8 + 12 = 23` and indices `[0,3,5]`. No other middle start produces a larger total. Although window start `1` ties window start `0` on the left, the left table retains `0`, which gives the lexicographically smaller answer.

## Why it works

For a fixed middle start `mid`, non-overlap requires the first window to start no later than `mid-k` and the third to start no earlier than `mid+k`. Maximizing those two windows independently maximizes their combined contribution because their allowed ranges do not depend on each other once `mid` is fixed. Therefore, the candidate formed from `left[mid-k]`, `mid`, and `right[mid+k]` is optimal for that middle position. Considering every valid `mid` includes every feasible triple, so the best candidate is globally optimal.

The left and right scans break equal-sum ties in favor of the earliest start. The middle scan proceeds from left to right and replaces the answer only on a strictly greater sum. Consequently, among maximum-sum triples, the algorithm retains the earliest first index, then the earliest middle index for that choice, and finally the earliest third index, yielding the lexicographically smallest answer.

## Complexity

Let `n = len(nums)`.

- **Time: `O(n)`.** Computing the window sums, both best-index tables, and the middle scan each take linear time.
- **Auxiliary space excluding the result: `O(n)`.** The window sums and left/right tables each contain `O(n)` entries.
- **Output space: `O(1)`.** The result always contains exactly three indices.

## Edge cases

[Tests](solution_test.go) cover both LeetCode examples, the minimum possible input length, equal values, separated maxima, ties between equal maxima, windows longer than one element, and the maximum input length. They also verify that the input is unchanged.

## Run

From the repository root:

```sh
go test -v ./problems/0689-maximum-sum-of-3-non-overlapping-subarrays
```

## Interview follow-ups

- Prefix sums can replace the sliding-window sum calculation with the same asymptotic complexity.
- A three-state dynamic program can reduce auxiliary storage, but reconstructing the indices and preserving lexicographic tie-breaking requires careful parent tracking.
- If only the maximum sum were required, the three starting indices and their reconstruction logic would not be needed.
