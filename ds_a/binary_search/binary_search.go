package binarysearch

func BinarySearchOrdered(arr []int, target int) int {
	// Establish upper and lower bounds

	// arr = [1, 2, 3, 4, 5, 6, 7, 8]
	// target = 7

	lowerBounds := 0
	upperBounds := len(arr) - 1 // 7

	for lowerBounds <= upperBounds {
		midpoint := (lowerBounds + upperBounds) / 2 // 3, 5, 6
		valueAtMidpoint := arr[midpoint]            // 4, 6, 7

		// If target is the midpoint, return index of midpoint
		if target == valueAtMidpoint {
			return midpoint
		} else if target < valueAtMidpoint {
			upperBounds = midpoint - 1
		} else if target > valueAtMidpoint {
			lowerBounds = midpoint + 1 // 4, 6
		}
	}

	return 0
}
