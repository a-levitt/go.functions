package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	tripleNumber := transformNumbers(&numbers, triple)
	fmt.Println("Triple:", tripleNumber)
	doubleNumber := transformNumbers(&numbers, double)
	fmt.Println("Double:", doubleNumber)
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
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
