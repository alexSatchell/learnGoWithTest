package linearsearch

func LinearSearchOrdered(arr []int, target int) int {
	for index, element := range arr {
		if element == target {
			return index
		} else if element > target {
			break
		}
	}
	return 0
}
