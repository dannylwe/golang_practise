package employee

// import "fmt"

// Employee struct definition
type Employee struct {
	FirstName, LastName string
}

// FullName method returns fullname of employee
func (e Employee) FullName() string {
	return e.FirstName + " " + e.LastName
}

// function changes first name. Without pointer, it creates copy
func (e *Employee) changeNameFirstName(newFirstName string) {
	e.FirstName = newFirstName
}

// func main() {
// 	employee := Employee{
// 		FirstName: "Daniel",
// 		LastName:  "Roy",
// 	}
// 	fmt.Println(employee.FullName())
// 	println("changing first name...\n")
// 	employee.changeNameFirstName("Danny Boy")
// 	fmt.Println(employee.FullName())
// }
