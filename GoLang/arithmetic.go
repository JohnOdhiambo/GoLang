package main

import (
	"fmt"
	"math"
)

func main() {
	var num1 float32 = 13
	var num2 float32 = 4

	var num3 int = 9
	var num4 int = 4

	ans := num1 / num2
	ans1 := num3 % num4
	ans2 := math.Sqrt(25)

	fmt.Println("The division is:", ans)    //The division is: 3.25
	fmt.Println("The modulus is:", ans1)    //The modulus is: 1
	fmt.Println("The squareroot is:", ans2) //The squareroot is: 5
}
