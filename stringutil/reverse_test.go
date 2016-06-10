package stringutil

import "testing"

func TestReverse(t *testing.T) {
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
			t.Errorf("Reverse(%v) == %v, want %v", c.s, got, c.want)
		}
	}
}

func BenchmarkReverse(b *testing.B) {
	for i := 1; i < b.N; i++ {
		Reverse("ReverseTest")
	}
}
