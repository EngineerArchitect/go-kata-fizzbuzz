package main

import (
	"fmt"
	"math"
	"strconv"
)

func isMultipleOfThree(x int) bool {
	return math.Mod(float64(x), 3) == 0
}

func isMultipleOfFive(x int) bool {
	return math.Mod(float64(x), 5) == 0
}

func fizzbuzz(x int) string {

	switch {
	case isMultipleOfThree(x) && isMultipleOfFive(x):
		return "FizzBuzz"
	case isMultipleOfThree(x):
		return "Fizz"
	case isMultipleOfFive(x):
		return "Buzz"
	default:
		return strconv.Itoa(x)
	}
}

func main() {
	for i := 1; i <= 15; i++ {
		fmt.Println(fizzbuzz(15))
	}
}
