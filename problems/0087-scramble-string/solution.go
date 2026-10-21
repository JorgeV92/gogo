package solution

// isScramble reports whether s2 can be obtained by recursively splitting s1
// and optionally swapping the two pieces. Inputs contain lowercase English letters.
func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}

	n := len(s1)
	// memo[i][j][length]: 0 = unknown, 1 = true, -1 = false.
	memo := make([][][]int8, n)
	for i := range memo {
		memo[i] = make([][]int8, n)
		for j := range memo[i] {
			memo[i][j] = make([]int8, n+1)
		}
	}

	var solve func(i, j, length int) bool
	solve = func(i, j, length int) bool {
		if memo[i][j][length] != 0 {
			return memo[i][j][length] == 1
		}
		if s1[i:i+length] == s2[j:j+length] {
			memo[i][j][length] = 1
			return true
		}

		// Splitting and swapping cannot change character frequencies.
		var counts [26]int
		for offset := 0; offset < length; offset++ {
			counts[s1[i+offset]-'a']++
			counts[s2[j+offset]-'a']--
		}
		for _, count := range counts {
			if count != 0 {
				memo[i][j][length] = -1
				return false
			}
		}

		for split := 1; split < length; split++ {
			// Keep the two pieces in the same order.
			if solve(i, j, split) && solve(i+split, j+split, length-split) {
				memo[i][j][length] = 1
				return true
			}
			// Match s1's left piece to s2's suffix, and its right to s2's prefix.
			if solve(i, j+length-split, split) && solve(i+split, j, length-split) {
				memo[i][j][length] = 1
				return true
			}
		}

		memo[i][j][length] = -1
		return false
	}

	return solve(0, 0, n)
}
