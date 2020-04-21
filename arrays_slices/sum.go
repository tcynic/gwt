package main

// Sum takes a slice of integers and returns their sum
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll takes a variable number of arrays
// and returns the sum of each array in an array
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}

// SumAllTails takes a variable number of arrays
// and returns the sum of the tail of each array in an array
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}
