package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	example := `{
        "id" : 11,
        "name" : "Irshad",
        "department" : "IT",
        "designation" : "Product Manager"
	}`

	var result map[string]interface{}

	json.Unmarshal([]byte(example), &result)

	// fmt.Println(result)
	fmt.Println(
		"Id:", result["id"],
		"\nName:", result["name"],
		"\nDepartment:", result["department"],
		"\nDesignation:", result["designation"],
	)
}
