package bubblesort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	t.Run("unsorted arrary", func(t *testing.T) {
		unsortedArr := []int{4, 2, 5, 1, 3}
		got := BubbleSort(unsortedArr)
		want := []int{1, 2, 3, 4, 5}

		// previously was attempting to use got != want
		// however you cannot compare to slices this way in go
		// use the reflect.DeepEqual function.
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
