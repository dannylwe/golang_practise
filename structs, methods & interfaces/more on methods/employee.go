package main

import "fmt"

type Employee struct {
	FirstName, LastName string
}

func FullName(firstName, lastName string) string {
	fullName := firstName + lastName
	return fullName
}
func main() {
	e := Employee{
		FirstName: "Daniel",
		LastName:  "Roy",
	}
	fmt.Println(FullName(e.FirstName, e.LastName))
}
