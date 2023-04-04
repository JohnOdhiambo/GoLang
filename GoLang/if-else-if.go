package main

import "fmt"

func main() {
	age := 22

	if age < 18 {
		fmt.Println("You are not eligible")
	} else if age > 18 && age < 25 {
		fmt.Printf("You are eligible for a limited time, Wait for %d year(s) more years", 25-age)
	} else {
		fmt.Printf("You are eligible to all levels")
	}
}
