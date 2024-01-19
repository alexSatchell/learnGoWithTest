package bubblesort

import "fmt"

// [4 2 5 1 3]
// [2 4 5 1 3]
// [2 4 5 1 3]
// [2 4 1 5 3]
// [2 4 1 3 5]
// [2 4 1 3 5]
// [2 1 4 3 5]
// [2 1 3 4 5]
// [2 1 3 4 5]
// [1 2 3 4 5]
// [1 2 3 4 5]
// [1 2 3 4 5]
// [1 2 3 4 5]
// [1 2 3 4 5]
// [1 2 3 4 5]
// [1 2 3 4 5]

// [4, 2, 5, 1, 3]
/*
* Pass through
* i = 0 [4 > 2] ? YES -> [2,4,5,1,3]
* i = 1 [5 > 2] ? NO ->  [2,4,5,1,3]
* i = 2 [5 > 1] ? YES -> [2,4,1,5,3]
* i = 3 [5 > 3] ? YES -> [2,4,1,3,5]
* i= 4
 */

func BubbleSort(arr []int) []int {
	// sorted := false
	//
	// for !sorted {
	// 	sorted = true
	// 	fmt.Println("started from Fresh")
	// 	for i := 0; i < len(arr)-1; i++ {
	// 		fmt.Println(arr)
	// 		if arr[i] > arr[i+1] {
	// 			arr[i], arr[i+1] = arr[i+1], arr[i]
	// 			sorted = false
	// 		}
	// 	}
	// }
	swapped := true
	for swapped {
		swapped = false
		fmt.Println("started from Fresh")
		for i := 1; i < len(arr); i++ {
			fmt.Println(arr)
			if arr[i-1] > arr[i] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
				swapped = true
			}
		}
	}
	return arr
}
