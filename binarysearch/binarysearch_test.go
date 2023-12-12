package binarysearch

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr1 := []int{1, 3, 4, 5, 9, 18, 20, 288}
	arr2 := []int{1, 3, 5, 9, 9, 19, 29, 29, 39}

	tests := []struct {
		data     []int
		target   int
		position int
		found    bool
	}{
		{arr1, 9, 4, true},
		{arr1, 388, 0, false},
		{arr2, 29, 7, true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("searching element %v", tt.target), func(t *testing.T) {
			pos, found := binarySearch(tt.data, tt.target)
			if pos != tt.position || found != tt.found {
				t.Errorf("got (%v, %v), want (%v, %v)", pos, found, tt.position, tt.found)
			}
		})
	}
}

func BenchmarkBinary_10(b *testing.B) {
	arr := generateArray(10)
	for i := 0; i < b.N; i++ {
		binarySearch(arr, 3)
	}
}

func BenchmarkBinary_100(b *testing.B) {
	arr := generateArray(100)
	for i := 0; i < b.N; i++ {
		binarySearch(arr, 43)
	}
}

func BenchmarkBinary_1000(b *testing.B) {
	arr := generateArray(1000)
	for i := 0; i < b.N; i++ {
		binarySearch(arr, 343)
	}
}

func BenchmarkBinary_10000(b *testing.B) {
	arr := generateArray(10000)
	for i := 0; i < b.N; i++ {
		binarySearch(arr, 54)
	}
}

func BenchmarkBinary_100000(b *testing.B) {
	arr := generateArray(100000)
	for i := 0; i < b.N; i++ {
		binarySearch(arr, 99999)
	}
}

func BenchmarkBinary_1000000(b *testing.B) {
	arr := generateArray(1000000)
	for i := 0; i < b.N; i++ {
		binarySearch(arr, 99923)
	}
}

func generateArray(numberElements int) []int {
	tmp := make([]int, numberElements)

	for i := 0; i < numberElements; i++ {
		tmp = append(tmp, i)
	}

	return tmp
}
