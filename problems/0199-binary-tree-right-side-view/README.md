# 199. Binary Tree Right Side View

| Field | Value |
| --- | --- |
| Problem | [LeetCode 199](https://leetcode.com/problems/binary-tree-right-side-view/) |
| Source | [Google question list](https://github.com/snehasishroy/leetcode-companywise-interview-questions/blob/master/google/all.csv) |
| Difficulty | Medium |
| Acceptance / frequency | 70.7% / 37.5% (supplied source snapshot) |
| Patterns | Binary tree, depth-first search |
| Status | Solved |

## Problem in my own words

Return the value of the rightmost node at each depth of a binary tree, starting at the root. An empty tree produces an empty result. A node in the left subtree can be visible when the right subtree does not reach that depth. See the [full statement and constraints](https://leetcode.com/problems/binary-tree-right-side-view/description/).

## Approach

A breadth-first traversal could process one level at a time and record its last node. Instead, this implementation uses depth-first search, visiting the right subtree before the left. At any given depth, the first node encountered is the rightmost node at that depth.

1. Start at the root with depth `0` and an empty result slice.
2. Return immediately for a `nil` node.
3. If `depth == len(view)`, this is the first node seen at this depth. Append its value.
4. Visit the right child and then the left child, each at `depth + 1`.

The result length counts depths already recorded. Later nodes at an existing depth are ignored, but their children are still explored because they may reach new depths.

[solution.go](solution.go) reads the tree without modifying its values or links. The result and recursive closure are local to each call.

## Walkthrough

Consider this tree:

```text
        8
       / \
      4   12
       \  /
        6 10
       /
      5
```

| Node visited | Depth | Action | Result |
| --- | --- | --- | --- |
| 8 | 0 | First at depth 0: append | `[8]` |
| 12 | 1 | First at depth 1: append | `[8, 12]` |
| 10 | 2 | First at depth 2: append | `[8, 12, 10]` |
| 4 | 1 | Depth already recorded | `[8, 12, 10]` |
| 6 | 2 | Depth already recorded | `[8, 12, 10]` |
| 5 | 3 | First at depth 3: append | `[8, 12, 10, 5]` |

The left child of `12` is visible at depth 2. The left subtree supplies `5` at depth 3 because no node in the right subtree reaches that depth. Following only right-child pointers would miss both nodes.

## Why it works

For two different nodes at the same depth, consider the ancestor where their paths first diverge. The node whose path enters that ancestor's right subtree is to the right of the other node. Since the traversal explores the entire right subtree before the left subtree, it encounters that node first. Thus the first node visited at every depth is the rightmost one.

The root is recorded first. Before reaching any deeper node, the traversal visits its ancestors, so recorded depths always form a consecutive prefix starting at zero. Therefore `depth == len(view)` identifies exactly the first visit to a new depth. Appending only on that visit records one correct value per depth in top-to-bottom order.

## Complexity

Let `n` be the number of nodes and `h` the number of levels in the tree.

- **Time: `O(n)`.** Every node is visited once; appending costs amortized constant time per visible node.
- **Auxiliary space excluding output: `O(h)`** for the recursion stack. This is `O(log n)` for a balanced tree and `O(n)` for a chain.
- **Output space: `O(h)`** for one value per level.

An empty tree takes constant time and space. These are asymptotic bounds, not measured LeetCode execution times.

## Edge cases

[Tests](solution_test.go) cover an empty tree, one node, complete and sparse trees, a deeper left subtree, competing nodes in opposite subtrees, a zigzag chain, repeated values, negative and boundary values, and a 100-node left chain. Results are compared in order; node values do not determine visibility.

## Run and submit

From the repository root:

```sh
go test -v ./problems/0199-binary-tree-right-side-view
go test ./...
```

Submit `rightSideView` from `solution.go`, omitting `package solution`. LeetCode supplies `TreeNode`; [tree_node.go](tree_node.go) exists for local compilation and should not be submitted. No imports are needed.

## Interview follow-ups

- A breadth-first solution takes `O(n)` time and `O(w)` auxiliary space, where `w` is maximum tree width, using a queue that retains only the current and next levels. Record the last node of each level when processing children left to right.
- An explicit stack avoids recursion. Push the left child before the right child so the right child is popped first, and store each node's depth.
- For a left-side view, visit the left subtree first and keep the same first-visit rule.
