# gogo

Go solutions and interview notes for problems in the [Google question list](https://github.com/snehasishroy/leetcode-companywise-interview-questions/blob/master/google/all.csv). Use that list as the backlog and the table below to track your work. Company tags and percentages come from this community-maintained source, not a guarantee of what Google will ask.

## Structure

```text
gogo/
├── go.mod
├── README.md                         # Setup, workflow, and progress
├── templates/
│   ├── README.md                      # Explanation template
│   ├── solution.go.tmpl               # Go starter
│   └── solution_test.go.tmpl          # Table-driven test starter
└── problems/
    └── 1471-the-k-strongest-values-in-an-array/
        ├── README.md                 # Approach, walkthrough, correctness, complexity
        ├── solution.go
        └── solution_test.go
```

Use `<four-digit-id>-<leetcode-slug>` for each problem directory. Each directory is an independent Go package named `solution`, so LeetCode function names can be reused across problems. A single root module runs all tests without external dependencies.

## Run locally

Requires Go 1.23 or newer. From the repository root:

```sh
# Test every implemented solution.
go test ./...

# Test one problem with individual case results.
go test -v ./problems/1471-the-k-strongest-values-in-an-array
```

## Progress

Percentages preserve the supplied source row; they are not live statistics. Add rows as you start problems. Suggested statuses: Not started, In progress, Solved, Revisit.

| ID | Problem / notes | Difficulty | Acceptance | Frequency | Patterns | Status |
| --- | --- | --- | --- | --- | --- | --- |
| 1471 | [The k Strongest Values in an Array](problems/1471-the-k-strongest-values-in-an-array/README.md) | Medium | 62.8% | 25.0% | Sorting, two pointers | Solved (example) |

## Add a problem

1. Pick a problem from the linked CSV and create its directory. For example:

   ```sh
   mkdir -p problems/0001-two-sum
   cp templates/README.md problems/0001-two-sum/README.md
   cp templates/solution.go.tmpl problems/0001-two-sum/solution.go
   cp templates/solution_test.go.tmpl problems/0001-two-sum/solution_test.go
   ```

2. Replace the starter function with LeetCode's Go signature and implement it. Keep helpers within the same package.
3. Update the test inputs, expected results, and function call. Cover examples, boundaries, and cases that distinguish your approach from an incorrect one. For unordered results, compare sorted copies or frequency maps.
4. Fill in the explanation template, including correctness, time complexity, auxiliary space, output space, and whether the input is modified.
5. Format and test the problem:

   ```sh
   gofmt -w problems/0001-two-sum/*.go
   go test -v ./problems/0001-two-sum
   go test ./...
   ```

6. Add a progress row with the original problem metadata and your status.

To submit on LeetCode, copy the imports, function, and any helpers from `solution.go`; omit the `package solution` line and the test file. Local solutions use `go test`, so no `main` function is needed.
