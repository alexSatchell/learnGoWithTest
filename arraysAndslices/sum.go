package main

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

// Variadic Params
// takes a variable length of ags of specified type
// note they will be combined into a slice
// here we have numbersToSum === []int [{1, 2}, {0, 9}]
func SumAll(numbersToSum ...[]int) []int {
	// How many items are in numbersToSum
	// 2 in the above example {1,2} & {0,9}
	lengthOfNumbers := len(numbersToSum)

	// Allocates & Inits a new slice, with a length of 2
	sums := make([]int, lengthOfNumbers)

	// For each slice e.g. {1,2}
	// in the numbersToSum slice we passed in
	// set our newly allocated slice at
	// the specified index equal to the sum
	// of that slice e.g. [3]
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
