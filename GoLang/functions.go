package main

import (
	"fmt"
)

func calcSum(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

// Recursive function-a function that calls itself
func countDown(number int) {
	if number > 0 {
		fmt.Println(number)
		//recursive call
		countDown(number - 1)
	} else {
		fmt.Println("Countdown has stopped")
	}
}

func main() {
	result := calcSum(4, 5)

	fmt.Println("The sum is", result)
	countDown(6)
}
