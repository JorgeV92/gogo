package solution

import (
	"strings"
	"testing"
)

func TestIsScramble(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want bool
	}{
		{name: "example nested swap", s1: "great", s2: "rgeat", want: true},
		{name: "example same counts but impossible", s1: "abcde", s2: "caebd", want: false},
		{name: "example single character", s1: "a", s2: "a", want: true},
		{name: "different single characters", s1: "a", s2: "b", want: false},
		{name: "identical strings", s1: "algorithm", s2: "algorithm", want: true},
		{name: "two character swap", s1: "ab", s2: "ba", want: true},
		{name: "unequal size swapped pieces", s1: "abc", s2: "bca", want: true},
		{name: "swaps in both children", s1: "great", s2: "rgtae", want: true},
		{name: "full reversal", s1: "abcdef", s2: "fedcba", want: true},
		{name: "duplicate letters", s1: "aabb", s2: "bbaa", want: true},
		{name: "same letter set different counts", s1: "aabb", s2: "abbb", want: false},
		{name: "interleaving is not always a scramble", s1: "abcd", s2: "bdac", want: false},
		{name: "maximum length equal", s1: strings.Repeat("a", 30), s2: strings.Repeat("a", 30), want: true},
		{name: "maximum length swapped halves", s1: strings.Repeat("a", 15) + strings.Repeat("b", 15), s2: strings.Repeat("b", 15) + strings.Repeat("a", 15), want: true},
		// Defensive behavior beyond LeetCode's nonempty, equal-length inputs.
		{name: "different lengths", s1: "ab", s2: "a", want: false},
		{name: "both empty", s1: "", s2: "", want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isScramble(tt.s1, tt.s2); got != tt.want {
				t.Fatalf("isScramble(%q, %q) = %v; want %v", tt.s1, tt.s2, got, tt.want)
			}
		})
	}
}
