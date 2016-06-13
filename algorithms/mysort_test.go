package algorithms

import (
	"math"
	"sort"
	"testing"
)

var (
	ints     = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	float64s = [...]float64{74.3, 59.0, math.Inf(1), 238.2, -784.0, 2.3, math.NaN(), math.NaN(), math.Inf(-1), 9845.768, -959.7485, 905, 7.8, 7.8}
	strings  = [...]string{"", "Hello", "foo", "bar", "foo", "f00", "%*&^*&^&", "***"}
	//runes = [...]rune{'世', 'a', 'z', 'd', '世', 'e', 'A', '界', 'c', 'b', '界', 'Q'}
	//bytes = [...]byte{'a', 'z', 'd', 'e', 'A', 'c', 'b', 'Q'}
)

func TestBubbleSort(t *testing.T) {
	intip := ints
	BubbleSort(sort.IntSlice(intip[0:]))
	if !sort.IsSorted(sort.IntSlice(intip[0:])) {
		t.Errorf("BubbleSort failed to sort: %v, result after sort %v", ints, intip)
	}

	stringip := strings
	BubbleSort(sort.StringSlice(stringip[0:]))
	if !sort.IsSorted(sort.StringSlice(stringip[0:])) {
		t.Errorf("BubbleSort failed to sort: %s, result after sort %s", strings, stringip)
	}

	float64ip := float64s
	BubbleSort(sort.Float64Slice(float64ip[0:]))
	if !sort.IsSorted(sort.Float64Slice(float64ip[0:])) {
		t.Errorf("BubbleSort failed to sort: %v, result after sort %v", float64s, float64ip)
	}
}

func TestSelectionSort(t *testing.T) {
	intip := ints
	SelectionSort(sort.IntSlice(intip[0:]))
	if !sort.IsSorted(sort.IntSlice(intip[0:])) {
		t.Errorf("SelectionSort failed to sort: %v, result after sort %v", ints, intip)
	}

	stringip := strings
	SelectionSort(sort.StringSlice(stringip[0:]))
	if !sort.IsSorted(sort.StringSlice(stringip[0:])) {
		t.Errorf("SelectionSort failed to sort: %s, result after sort %s", strings, stringip)
	}

	float64ip := float64s
	SelectionSort(sort.Float64Slice(float64ip[0:]))
	if !sort.IsSorted(sort.Float64Slice(float64ip[0:])) {
		t.Errorf("SelectionSort failed to sort: %v, result after sort %v", float64s, float64ip)
	}
}

func TestInsertionSort(t *testing.T) {
	intip := ints
	InsertionSort(sort.IntSlice(intip[0:]))
	if !sort.IsSorted(sort.IntSlice(intip[0:])) {
		t.Errorf("InsertionSort failed to sort: %v, result after sort %v", ints, intip)
	}

	stringip := strings
	InsertionSort(sort.StringSlice(stringip[0:]))
	if !sort.IsSorted(sort.StringSlice(stringip[0:])) {
		t.Errorf("InsertionSort failed to sort: %s, result after sort %s", strings, stringip)
	}

	float64ip := float64s
	InsertionSort(sort.Float64Slice(float64ip[0:]))
	if !sort.IsSorted(sort.Float64Slice(float64ip[0:])) {
		t.Errorf("InsertionSort failed to sort: %v, result after sort %v", float64s, float64ip)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	intip := ints
	for i := 1; i < b.N; i++ {
		BubbleSort(sort.IntSlice(intip[0:]))
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	intip := ints
	for i := 1; i < b.N; i++ {
		SelectionSort(sort.IntSlice(intip[0:]))
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	intip := ints
	for i := 1; i < b.N; i++ {
		InsertionSort(sort.IntSlice(intip[0:]))
	}
}
