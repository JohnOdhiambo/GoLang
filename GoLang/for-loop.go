package main

import "fmt"

func main() {

	for x := 0; x <= 30; x++ {
		if x != 0 && x%3 == 0 && x%5 == 0 {
			fmt.Println(x)
			break
		}
	}
}
