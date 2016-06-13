package stringutil

import (
	"math/rand"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	// test data with expected result
	var tests = []struct {
		s            string
		isPalindrome bool
	}{
		{"", false},
		{"m", true},
		{"ace", false},
		{"ana", true},
		{"anna", true},
		{"aloha", false},
		{"alola", true},
		{"civic", true},
		{"toyota", false},
		{"not a palindrome", false},
	}

	// for each entry in test data apply function and compare with expected result
	for _, c := range tests {
		got := isPalindrome(c.s)
		if got != c.isPalindrome {
			// fail if the result doesn't match expected
			t.Errorf("isPalindrome(%v) == %v, want %v", c.s, got, c.isPalindrome)
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	// input array
	ip := []string{"", "civic", "a", "aloha", "ahoha", "palindrome", "verylongstringforpalindrometest", "aaaaabbbbcccddeddcccbbbbaaaaa"}
	for i := 1; i < b.N; i++ {
		// random input selection from the given input array
		isPalindrome(ip[rand.Intn(len(ip))])
	}
}
