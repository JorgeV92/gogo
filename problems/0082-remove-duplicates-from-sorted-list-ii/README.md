# 82. Remove Duplicates from Sorted List II

| Field | Value |
| --- | --- |
| Problem | [LeetCode 82](https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/) |
| Source | [Google question list](https://github.com/snehasishroy/leetcode-companywise-interview-questions/blob/master/google/all.csv) |
| Difficulty | Medium |
| Acceptance / frequency | 52.1% / 25.0% (supplied source snapshot) |
| Patterns | Linked list, dummy head, two pointers |
| Status | Solved |

## Problem in my own words

Given a sorted singly linked list, remove every node whose value appears more than once. Only values that appeared exactly once survive, in their original order. For example, `[1, 2, 2, 4]` becomes `[1, 4]`. An empty list is valid. See the [full statement and constraints](https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/description/).

## Approach

A frequency map could count occurrences before a second pass removes repeated values, using `O(n)` extra space. Since the input is sorted, equal values are adjacent: a single pass can process whole runs with constant extra space.

Use a dummy node before the head so removing the first run works just like removing any later run. Its value is irrelevant; only its `Next` pointer is used.

1. Set `prev` to the dummy node and `current` to the head. `prev` is the last retained node, or the dummy if none have been retained.
2. If `current` and its successor have equal values, save that value and advance `current` past every node in that run. Set `prev.Next = current` to unlink the whole run. Keep `prev` in place.
3. Otherwise, retain the unique node by advancing `prev` to it and `current` to its successor.
4. Return `dummy.Next`, which may differ from the original head or be `nil`.

[solution.go](solution.go) modifies the input's links and reuses all surviving nodes. The caller should use the returned head. It does not allocate a new result list.

## Walkthrough

For `[1, 2, 2, 3, 4, 4, 5]`:

| Run | Action | `prev` afterward | List reachable from dummy afterward |
| --- | --- | --- | --- |
| `1` | Retain unique node | `1` | `[1, 2, 2, 3, 4, 4, 5]` |
| `2, 2` | Link `1` directly to `3` | `1` | `[1, 3, 4, 4, 5]` |
| `3` | Retain unique node | `3` | `[1, 3, 4, 4, 5]` |
| `4, 4` | Link `3` directly to `5` | `3` | `[1, 3, 5]` |
| `5` | Retain unique node | `5` | `[1, 3, 5]` |

For `[2, 2, 3]`, the first run is removed by changing the dummy's `Next` to the node containing `3`. No special head-removal branch is needed.

## Why it works

At the start of each outer iteration, the processed prefix reachable from the dummy contains exactly the unique values from the processed runs. `prev` ends that prefix, and `prev.Next == current` points to the first unprocessed run.

Sorted order means a value cannot reappear after its run. A run of length one must therefore be retained; advancing both pointers extends the correct prefix. A longer run must be removed entirely; linking `prev` to the next run preserves the prefix and skips all occurrences. Both branches preserve the invariant and advance `current`.

When `current` reaches `nil`, all runs have been processed, so the returned list contains exactly the values that occurred once and remains sorted.

## Complexity

Let `n` be the number of input nodes and `u` the number of surviving nodes.

- **Time: `O(n)`.** `current` only moves forward. Across all iterations, including the inner loop, each node is passed once.
- **Auxiliary space: `O(1)`.** A dummy node, a few pointers, and one saved value; no recursion or map.
- **New output space: `O(1)`.** The returned pointer refers to `u` existing nodes; no new output nodes are allocated.

These are asymptotic bounds, not measured LeetCode execution times.

## Edge cases

[Tests](solution_test.go) cover an empty list, a single node, all distinct values, two equal nodes, longer duplicate runs, duplicates at the head or tail, consecutive duplicate runs, all values removed, only one middle node surviving, and negative/zero/boundary values. They verify output order, termination, and reuse of original nodes.

## Run and submit

From the repository root:

```sh
go test -v ./problems/0082-remove-duplicates-from-sorted-list-ii
go test ./...
```

On LeetCode, submit the function from `solution.go`, omitting `package solution`. LeetCode supplies `ListNode`; [list_node.go](list_node.go) exists only to compile locally and should not be submitted.

## Interview follow-ups

- For an unsorted list, count values with a map before unlinking repeated values: `O(n)` time and `O(n)` extra space.
- To preserve the original list, scan the runs and copy only singleton nodes into a new list, using `O(u)` new output space.
- For a variant that keeps one occurrence of each value, retain the first node in each run and skip the rest. This problem requires removing the entire repeated run.
