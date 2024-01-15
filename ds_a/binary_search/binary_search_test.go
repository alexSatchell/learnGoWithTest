package binarysearch

import "testing"

func TestBinarySearchOrdered(t *testing.T) {
	t.Run("Binary search on ordered array", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5, 6, 7, 8}

		got := BinarySearchOrdered(arr, 7)
		want := 6

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
