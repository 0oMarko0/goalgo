// Package binarysearch will return the position of an element
// in a sorted array in O(log n).
package binarysearch

func binarySearch(arr []int, elem int) (i int, found bool) {
	low := 0
	high := len(arr)

	for low <= high {
		mid := (high + low) / 2

		if mid < len(arr) && arr[mid] == elem {
			return mid, true
		}

		if elem > mid {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return 0, false
}
