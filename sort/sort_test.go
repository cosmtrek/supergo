package sort

import (
	"math/rand"
	s "sort"
	"testing"
)

func TestInsertSort(t *testing.T) {
	arr := []int{1, 3, 6, 9, 2, 4, 8, 0, 5, 7}
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	InsertSort(arr)
	if i, ok := compareArray(arr, want); !ok {
		t.Errorf("Got: %v, want: %v, index: %d, got: %d, want: %d", arr, want, i, arr[i], want[i])
	}
}

func TestHeapSort(t *testing.T) {
	arr := []int{1, 3, 6, 9, 2, 4, 8, 0, 5, 7}
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	HeapSort(arr)
	if i, ok := compareArray(arr, want); !ok {
		t.Errorf("Got: %v, want: %v, index: %d, got: %d, want: %d", arr, want, i, arr[i], want[i])
	}
}

func TestHeapSortBigArray(t *testing.T) {
	var arr []int
	var want []int

	for i := 0; i < 1000000; i++ {
		n := rand.Intn(100)
		arr = append(arr, n)
		want = append(want, n)
	}

	s.IntSlice(want).Sort()
	HeapSort(arr)
	if i, ok := compareArray(arr, want); !ok {
		t.Errorf("index: %d, got: %d, want: %d", i, arr[i], want[i])
	}
}

func TestQuickSort(t *testing.T) {
	arr := []int{1, 3, 6, 9, 2, 4, 8, 0, 5, 7}
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	QuickSort(arr)
	if i, ok := compareArray(arr, want); !ok {
		t.Errorf("Got: %v, want: %v, index: %d, got: %d, want: %d", arr, want, i, arr[i], want[i])
	}
}

func TestQuickSortBigArray(t *testing.T) {
	var arr []int
	var want []int

	for i := 0; i < 1000000; i++ {
		n := rand.Intn(100)
		arr = append(arr, n)
		want = append(want, n)
	}

	s.IntSlice(want).Sort()
	QuickSort(arr)
	if i, ok := compareArray(arr, want); !ok {
		t.Errorf("index: %d, got: %d, want: %d", i, arr[i], want[i])
	}
}

func compareArray(got []int, want []int) (int, bool) {
	for i := range got {
		if got[i] != want[i] {
			return i, false
		}
	}
	return -1, true
}

func TestShuffle(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	Shuffle(arr)
}
