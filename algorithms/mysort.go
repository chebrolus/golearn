package algorithms

import "sort"

/*
 * Good old bubble sort to start with: https://en.wikipedia.org/wiki/Bubble_sort
 * Using golang sort.Interface to be able to re-use this function for various data types
 * sort.Interface currently support int, string & float64.
 */
func BubbleSort(ip sort.Interface) {
	swapped := true
	length := ip.Len()

	// outer loop, reduce length as each iteration end up moving largest # to right most
	for ; swapped; length-- {
		// reset swapped flag here
		swapped = false
		for i := 1; i < length; i++ {
			// ith location value < i-1 location value then swap and set swapped flag
			if ip.Less(i, i-1) {
				ip.Swap(i, i-1)
				swapped = true
			}
		}
	}
}

// https://en.wikipedia.org/wiki/Selection_sort
func SelectionSort(ip sort.Interface) {
	var min int
	length := ip.Len()
	for i := 0; i < length; i++ {
		min = i
		for j := i + 1; j < length; j++ {
			if ip.Less(j, min) {
				min = j
			}
		}
		ip.Swap(i, min)
	}
}

// https://en.wikipedia.org/wiki/Insertion_sort
func InsertionSort(ip sort.Interface) {
	length := ip.Len()

	for i := 1; i < length; i++ {
		for j := i; j > 0 && ip.Less(j, j-1); j-- {
			ip.Swap(j, j-1)
		}
	}
}
