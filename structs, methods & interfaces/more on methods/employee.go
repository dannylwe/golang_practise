package main

import "fmt"

type Employee struct {
	FirstName, LastName string
}

func (e Employee) FullName() string {
	return e.FirstName + " " + e.LastName
}
func main() {
	employee := Employee{
		FirstName: "Daniel",
		LastName:  "Roy",
	}
	fmt.Println(employee.FullName())
}
