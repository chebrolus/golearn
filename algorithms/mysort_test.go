package algorithms

import (
	"testing"
	"sort"
	"math"
)

var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
var float64s = [...]float64{74.3, 59.0, math.Inf(1), 238.2, -784.0, 2.3, math.NaN(), math.NaN(), math.Inf(-1), 9845.768, -959.7485, 905, 7.8, 7.8}
var strings = [...]string{"", "Hello", "foo", "bar", "foo", "f00", "%*&^*&^&", "***"}

func TestBubbleSort(t *testing.T)  {
	intip := ints
	BubbleSort(sort.IntSlice(intip[0:]))
	if !sort.IsSorted(sort.IntSlice(intip[0:])) {
		t.Errorf("BubbleSort failed to sort: %s, result after sort %s", ints, intip)
	}

	stringip := strings
	BubbleSort(sort.StringSlice(stringip[0:]))
	if !sort.IsSorted(sort.StringSlice(stringip[0:])) {
		t.Errorf("BubbleSort failed to sort: %s, result after sort %s", strings, stringip)
	}

	float64ip := float64s
	BubbleSort(sort.Float64Slice(float64ip[0:]))
	if !sort.IsSorted(sort.Float64Slice(float64ip[0:])) {
		t.Errorf("BubbleSort failed to sort: %s, result after sort %s", float64s, float64ip)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	intip := ints
	for i := 1; i < b.N; i++ {
		BubbleSort(sort.IntSlice(intip[0:]))
	}
}
