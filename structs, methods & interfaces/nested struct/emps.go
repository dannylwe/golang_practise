package main

import "fmt"

type Contact struct {
	phone, address string
}

type Employee struct {
	name    string
	salary  int
	contact Contact
}

func (e *Employee) changePhone(newPhone string) {
	e.contact.phone = newPhone
}

func main() {
	employee := Employee{
		name:   "Mr. Fantastic",
		salary: 1000,
		contact: Contact{
			phone:   "234567",
			address: "fifwjn rd, kjhfjwen",
		},
	}
	fmt.Println("before phone change\n", employee)
	println("changing phone number...\n")
	employee.changePhone("00000000")
	fmt.Println("after phone change", employee)
}
