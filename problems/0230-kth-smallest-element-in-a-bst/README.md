# 230. Kth Smallest Element in a BST

| Field | Value |
| --- | --- |
| Problem | [LeetCode 230](https://leetcode.com/problems/kth-smallest-element-in-a-bst/) |
| Source | [Google question list](https://github.com/snehasishroy/leetcode-companywise-interview-questions/blob/master/google/all.csv) |
| Difficulty | Medium |
| Acceptance / frequency | 77.0% / 37.5% (linked source snapshot) |
| Patterns | Binary search tree, inorder traversal, stack |
| Status | Solved |

## Problem in my own words

Find the value at rank `k` when a binary search tree's values are ordered from smallest to largest. Ranks start at one. The tree has `1 <= n <= 10,000` nodes, `1 <= k <= n`, and values from `0` through `10,000`. See the [full statement and constraints](https://leetcode.com/problems/kth-smallest-element-in-a-bst/description/).

## Approach

An inorder traversal visits a BST in sorted order: left subtree, node, then right subtree. Collecting all values into a slice would take `O(n)` time and space. We only need the kth visit, so an explicit stack lets us stop as soon as we reach it.

1. Push the current node and follow left children until reaching `nil`.
2. Pop the stack. This node holds the next smallest value.
3. Decrement `k`. If it becomes zero, return this node's value.
4. Move to the popped node's right child and repeat.

[solution.go](solution.go) reads the tree without changing values or links. Each call has its own stack and remaining rank. An empty tree or an out-of-range rank is outside the problem's contract; the final panic is unreachable for valid input.

## Walkthrough

For this tree, find rank `4`:

```text
        8
       / \
      3   10
       \
        6
       / \
      4   7
```

The stack below is shown from bottom to top, just before each pop.

| Stack | Value visited | Remaining rank | Next action |
| --- | --- | --- | --- |
| `[8, 3]` | 3 | 3 | Enter the right subtree rooted at 6 |
| `[8, 6, 4]` | 4 | 2 | No right child; pop the next ancestor |
| `[8, 6]` | 6 | 1 | Enter the right subtree rooted at 7 |
| `[8, 7]` | 7 | 0 | Return 7 |

The traversal returns before visiting 8 or 10. Rank counts nodes visited, so gaps between values do not affect it.

## Why it works

In a BST, values in a node's left subtree precede its value in sorted order, and values in its right subtree follow it. Applying that rule recursively shows that inorder traversal visits every value in increasing order.

The stack preserves ancestors whose own visits are pending. Following left children reaches the smallest unvisited value in the current subtree. Popping visits that value, and entering its right subtree processes the next values before returning to larger pending ancestors. Thus the iterative traversal preserves inorder order without skipping or revisiting nodes.

After `j` pops, exactly the `j` smallest values have been visited, and the remaining rank equals the original `k - j`. It reaches zero exactly at the requested value. A valid rank guarantees this occurs before traversal finishes.

## Complexity

Let `n` be the node count, `h` the tree height in nodes, and `k` the original requested rank.

- **Time: `O(h + k)`, and at most `O(n)`.** There are `k` pops. Each pushed node is either popped or remains on the stack, whose size is at most `h`. Slice appends take amortized constant time.
- **Auxiliary space excluding output: `O(h)`** for the stack: `O(log n)` for a balanced tree, `O(n)` in the worst case.
- **Output space: `O(1)`** for the returned integer.

## Edge cases

[Tests](solution_test.go) cover both statement examples, a single zero-valued node, first and last ranks, the root's rank, gaps between values, both value boundaries, sparse subtrees requiring left descent after a right turn, and 10,000-node left and right chains. Shared fixtures are queried at multiple ranks to exercise independent calls.

## Run and submit

From the repository root:

```sh
go test -v ./problems/0230-kth-smallest-element-in-a-bst
go test ./...
```

Submit `kthSmallest` from `solution.go`, omitting `package solution`. LeetCode supplies `TreeNode`; [tree_node.go](tree_node.go) supports local compilation and should not be submitted. No imports are needed.

## Interview follow-ups

- Recursive inorder traversal has the same asymptotic bounds if it propagates an early stop, but uses the call stack.
- For frequent rank queries with insertions and deletions, store each node's subtree size. If the left subtree has `L` nodes, compare `k` with `L + 1`: go left, return the current value, or go right with rank `k - L - 1`. Maintain sizes along update paths and during rotations. A balanced tree gives `O(log n)` queries and updates; an unbalanced tree can still require `O(n)`.
- Morris traversal can use constant auxiliary space by temporarily linking predecessors to ancestors. Returning early requires care to restore all temporary links; the explicit stack keeps the input untouched throughout.
