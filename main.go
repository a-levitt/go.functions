package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	transformedNumber := transformNumbers(&numbers, triple)

	fmt.Println(transformedNumber)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	trNumbers := []int{}
	for _, val := range *numbers {
		trNumbers = append(trNumbers, transform(val))
	}

	return trNumbers
}

func double(number int) int {
	return number * 2
}
func triple(number int) int {
	return number * 3
}
