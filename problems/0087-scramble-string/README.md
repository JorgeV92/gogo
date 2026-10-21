# 87. Scramble String

| Field | Value |
| --- | --- |
| Problem | [LeetCode 87](https://leetcode.com/problems/scramble-string/) |
| Source | [Google question list](https://github.com/snehasishroy/leetcode-companywise-interview-questions/blob/master/google/all.csv) |
| Difficulty | Hard |
| Acceptance / frequency | 45.0% / 25.0% (supplied source snapshot) |
| Patterns | Dynamic programming, memoized recursion |
| Status | Solved |

## Problem in my own words

A scramble recursively divides a string into two nonempty pieces, optionally exchanges their order, and repeats within each piece. Determine whether these operations can transform `s1` into `s2`. The inputs have equal lengths from 1 through 30 and contain lowercase English letters. See the [full statement](https://leetcode.com/problems/scramble-string/description/).

Matching character counts is necessary but not sufficient: `abcde` and `caebd` are anagrams, yet no valid sequence of recursive splits produces that transformation.

## Approach

Trying every split recursively repeats many substring comparisons. Memoize each comparison using three integers:

```text
solve(i, j, length)
```

This asks whether `s1[i:i+length]` can scramble into `s2[j:j+length]`. Go slice bounds here exclude the ending index. Both substrings have the same length, so three coordinates describe the entire state.

1. Return a cached answer if the state has already been evaluated.
2. If the substrings are identical, return true: keeping the order at every split works.
3. Count the letters in both substrings. If any counts differ, return false immediately.
4. For every split position `k` from `1` through `length-1`, try both arrangements:

   | Arrangement | First comparison | Second comparison |
   | --- | --- | --- |
   | Keep order | `solve(i, j, k)` | `solve(i+k, j+k, length-k)` |
   | Swap pieces | `solve(i, j+length-k, k)` | `solve(i+k, j, length-k)` |

   Both comparisons in an arrangement must succeed. Return true as soon as either arrangement works.

5. If no split works, return false. Cache both successful and failed states.

For a swapped split, the left piece of `s1` has length `k`, so it must match the last `k` characters of the target, starting at `j+length-k`. This offset also handles unequal piece sizes.

[solution.go](solution.go) uses a three-dimensional `int8` table: `0` means unknown, `1` means true, and `-1` means false. Distinguishing unknown from false prevents repeatedly evaluating impossible states. All memoization belongs to one function call. The strings are not modified, and Go substring expressions do not copy their bytes.

## Walkthrough

For `s1 = "abcd"` and `s2 = "dcab"`, try a root split of length 2:

```text
Source: ab | cd
Target: dc | ab
```

The swapped arrangement compares `ab` with `ab` (already equal) and `cd` with `dc`. The latter succeeds by splitting into individual characters and swapping them. Thus the root state returns true. A failed earlier split does not stop the search; only a successful arrangement does.

For a one-character state, equal letters return true in the equality check; different letters return false in the frequency check. No further split is possible or needed.

## Why it works

Use induction on substring length. For length one, the equality and frequency checks give the correct answer.

For longer substrings, a valid scramble has some first split, followed by either keeping or swapping its pieces. The algorithm tries every possible first split and both arrangements. By induction, each recursive comparison correctly decides whether the corresponding pieces can be transformed. Therefore a successful pair of comparisons constructs a valid scramble, and every valid scramble is considered by some pair. If no pair succeeds, the transformation is impossible.

The pruning checks are sound: identical strings can keep their order, and splitting or swapping never changes character counts. Memoization reuses answers without changing this reasoning.

## Complexity

Let `n` be the common string length.

- **Time: `O(n⁴)` worst case.** There are `O(n³)` states, one for each pair of starting positions and substring length. Each state performs `O(n)` work for equality, counting, and trying splits. Recursive calls reuse cached states. Pruning often avoids much of this work.
- **Auxiliary space: `O(n³)`.** The memo table dominates the `O(n)` recursion stack. Each stack frame has a fixed 26-element frequency array. The table is allocated eagerly once unequal strings of equal length reach the recursive solver.
- **Output space: `O(1)`.** The result is a boolean.

These are asymptotic bounds, not measured LeetCode execution times.

## Edge cases

[Tests](solution_test.go) cover single characters, identical strings, swapped pieces of unequal sizes, nested swaps, reversal, duplicate letters, unequal character counts, anagrams that cannot be scrambled, and the maximum allowed length. Defensive checks also return false for unequal lengths and true for two empty strings; these inputs are outside the stated constraints. The implementation assumes lowercase English letters.

## Run and submit

From the repository root:

```sh
go test -v ./problems/0087-scramble-string
go test ./...
```

Submit `isScramble` from `solution.go` to LeetCode, omitting the `package solution` line. No imports or external helpers are required.

## Interview follow-ups

- A bottom-up dynamic program fills the same states in increasing substring length, with the same worst-case bounds and no recursion stack.
- Prefix frequency tables make frequency comparisons `O(26)` per state, although trying all split positions still leaves the overall worst-case time at `O(n⁴)`.
- To return a sequence of swaps, store the successful split and arrangement for each true state and reconstruct a transformation tree.
