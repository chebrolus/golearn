package stringutil

import "testing"

func TestIsPalindrome(t *testing.T) {
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
	for _, c := range tests {
		got := isPalindrome(c.s)
		if got != c.isPalindrome {
			t.Errorf("isPalindrome(%v) == %v, want %v", c.s, got, c.isPalindrome)
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 1; i < b.N; i++ {
		isPalindrome("ReverseeveR")
	}
}
