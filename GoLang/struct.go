// Struct is a user defined type which represents a collection of fields
package main

import "fmt"

type Student struct {
	name   string
	age    int
	course string
}

func main() {
	//Creating a struct
	stud := Student{"Jay", 24, "IT"}

	fmt.Println("The student full info is:", stud)
	fmt.Println("Name:", stud.name)
	fmt.Println("Age:", stud.age)
	fmt.Println("Course:", stud.course)
}
