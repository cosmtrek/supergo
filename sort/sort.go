// Package sort aims to practice, but not ready for production :)
package sort

import (
	"math"
	"math/rand"
	"time"
)

// InsertSort sorts int array incrementally.
func InsertSort(arr []int) {
	if len(arr) == 0 {
		return
	}

	for i := range arr {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
}

// HeapSort sorts int array incrementally.
func HeapSort(arr []int) {
	heapSize := len(arr)
	buildHeap(arr, heapSize)

	for i := heapSize - 1; i >= 0; i-- {
		// place the max num in the last position
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, 0, i)
	}
}

func buildHeap(arr []int, heapSize int) {
	for i := int(math.Floor(float64(heapSize/2)) - 1); i >= 0; i-- {
		heapify(arr, i, heapSize)
	}
}

func heapify(arr []int, index int, heapSize int) {
	left, right, largest := 2*index+1, 2*index+2, index
	if left < heapSize && arr[index] < arr[left] {
		largest = left
	}
	if right < heapSize && arr[largest] < arr[right] {
		largest = right
	}
	if largest != index {
		arr[largest], arr[index] = arr[index], arr[largest]
		heapify(arr, largest, heapSize)
	}
}

// QuickSort sorts int array incrementally.
// reference: https://en.wikipedia.org/wiki/Quicksort
func QuickSort(arr []int) {
	if len(arr) == 0 {
		return
	}

	quickSortHelper(arr, 0, len(arr)-1)
}

func quickSortHelper(arr []int, left int, right int) {
	threshold := 20
	if left < right {
		lt, gt := partition(arr, left, right)
		if lt-1-left <= threshold {
			insertSortHelper(arr, left, lt-1)
		} else {
			quickSortHelper(arr, left, lt-1)
		}
		if right-gt-1 <= threshold {
			insertSortHelper(arr, gt+1, right)
		} else {
			quickSortHelper(arr, gt+1, right)
		}
	}
}

// 3way qsort
// reference: http://x-wei.github.io/quick-sort-and-more.html
// invariant:
//   a[lo,lt-1] < pivot
//   a[lt, i-1] = pivot
//   a[i,gt] = unseen
//   a[gt+1, hi] > pivot
func partition(arr []int, left int, right int) (int, int) {
	m := left + (right-left)/2
	if right-left > 40 {
		s := (right - left) / 8
		medianOfThree(arr, left, left+s, left+2*s)
		medianOfThree(arr, m, m-s, m+s)
		medianOfThree(arr, right-1, right-1-s, right-1-2*s)
	}
	medianOfThree(arr, left, m, right-1)

	lt, i, gt, pivot := left, left, right, arr[left]

	for i <= gt {
		if arr[i] == pivot {
			i++
		} else if arr[i] > pivot {
			arr[gt], arr[i] = arr[i], arr[gt]
			gt--
		} else {
			arr[lt], arr[i] = arr[i], arr[lt]
			lt++
			i++
		}
	}

	return lt, gt
}

func insertSortHelper(arr []int, left int, right int) {
	for i := left; i <= right; i++ {
		for j := i; j > left && arr[j-1] > arr[j]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
}

// Find the median of three numbers, and place the median in the most left place.
func medianOfThree(arr []int, left, mid, right int) {
	if arr[left] > arr[mid] {
		arr[left], arr[mid] = arr[mid], arr[left]
	}
	if arr[right] < arr[mid] {
		arr[mid], arr[right] = arr[right], arr[mid]
		if arr[left] > arr[mid] {
			arr[left], arr[mid] = arr[mid], arr[left]
		}
	}

	arr[left], arr[mid] = arr[mid], arr[left]
}

// Lomuto partition scheme, low efficiency.
// func partition(arr []int, left int, right int) int {
//	pivot := arr[right]
//	i := left
//
//	for j := left; j < right; j++ {
//		if arr[j] <= pivot {
//			arr[i], arr[j] = arr[j], arr[i]
//			i++
//		}
//	}
//	arr[i], arr[right] = arr[right], arr[i]
//
//	return i
// }

// Shuffle of Fisher-Yates https://en.wikipedia.org/wiki/Fisher–Yates_shuffle
func Shuffle(arr []int) []int {
	arrLen := len(arr)
	if arrLen == 0 || arrLen == 1 {
		return arr
	}

	rand.Seed(time.Now().UnixNano())
	for i := arrLen - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}
