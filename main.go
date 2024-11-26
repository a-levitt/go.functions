package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2}

	tripleNumber := transformNumbers(&numbers, triple)
	fmt.Println("Triple:", tripleNumber)
	doubleNumber := transformNumbers(&numbers, double)
	fmt.Println("Double:", doubleNumber)

	transformerFn1 := getTransformerFunction(&numbers)
	transformerFn2 := getTransformerFunction(&moreNumbers)
	transformedNumbers := transformNumbers(&numbers, transformerFn1)
	moreTransformedNumbers := transformNumbers(&moreNumbers, transformerFn2)
	fmt.Println(transformedNumbers, "|||", moreTransformedNumbers)
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	trNumbers := []int{}
	for _, val := range *numbers {
		trNumbers = append(trNumbers, transform(val))
	}

	return trNumbers
}

// func getTransformerFunction() func(int) int {
func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}
func triple(number int) int {
	return number * 3
}
