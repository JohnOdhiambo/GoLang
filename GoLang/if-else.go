package main

import "fmt"

func main() {
	age := 15

	if age > 18 {
		fmt.Println("You are eligible")
	} else {
		fmt.Println("You are not eligible")
		fmt.Printf("Wait for %d year(s)", 18-age)
	}
}
