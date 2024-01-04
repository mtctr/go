package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {

	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {

		tail := GetTail(numbers)
		sums = append(sums, Sum(tail))
	}
	return sums
}

func GetTail(numbers []int) []int {
	if len(numbers) == 0 {
		return []int{0, 0}
	}
	return numbers[1:]

}
