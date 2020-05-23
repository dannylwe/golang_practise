package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	array := `[
		{
			"id": 1,
			"name": "Mr. Boss",
			"department": "",
			"designation": "Director"
		},
		{
			"id": 11,
			"name": "Irshad",
			"department": "IT",
			"designation": "Product Manager"
		},
		{
			"id": 12,
			"name": "Pankaj",
			"department": "IT",
			"designation": "Team Lead"
		}
	]`
	var results []map[string]interface{}

	json.Unmarshal([]byte(array), &results)

	for key, val := range results {
		fmt.Println("reading key:", key)
		fmt.Println(
			"Id:", val["id"],
			"\nName:", val["name"],
			"\nDepartment:", val["department"],
			"\nDesignation:", val["designation"], "\n",
		)
	}
}
