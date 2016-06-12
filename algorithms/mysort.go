package algorithms

import "sort"

// Good old bubble sort to start with
// Using golang sort.Interface to be able to re-use this function for various data types
// sort.Interface currently support int, string & float64. Easy to extend this approach to other types
func BubbleSort(ip sort.Interface)  {
	swapped := true;
	length := ip.Len();

	// outer loop, reduce length as each iteration end up moving largest # to right most
	for ; swapped ; length-- {
		// reset swapped flag here
		swapped = false;
		for i := 1; i < length; i++ {
			// ith location value < i-1 location value then swap and set swapped flag
			if ip.Less(i, i - 1) {
				ip.Swap(i, i - 1);
				swapped = true;
			}
		}
	}
}