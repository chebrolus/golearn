package stringutil

import (
	"testing"
	"math/rand"
)

func TestReverse(t *testing.T) {
	// test data with expected result
	var tests = []struct {
		s, want string
	}{
		{"", ""},
		{"Hello, 世界", "界世 ,olleH"},
		{"Backward", "drawkcaB"},
	}
	for _, c := range tests {
		got := Reverse(c.s)
		if got != c.want {
			// fail if function response doesn't match with expected result
			t.Errorf("Reverse(%v) == %v, want %v", c.s, got, c.want)
		}
	}
}

func BenchmarkReverse(b *testing.B) {
	// input array
	ip := []string{"", "a", "ReverseTest", "verylongstringforreversetest", "verylongstringforreversetestverylongstringforreversetestverylongstringforreversetest"}
	for i := 1; i < b.N; i++ {
		// random input selection from the given input array
		Reverse(ip[rand.Intn(len(ip))])
	}
}
