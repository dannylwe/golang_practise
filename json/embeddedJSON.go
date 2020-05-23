package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	array := `{
		"id": 11,
		"name": "Irshad",
		"department": "IT",
		"designation": "Product Manager",
		"address": {
			"city": "Mumbai",
			"state": "Maharashtra",
			"country": "India"
		}
	}`

	var result map[string]interface{}

	json.Unmarshal([]byte(array), &result)
	
	// fmt.Println(result)
	address := result["address"].(map[string]interface{})

	fmt.Println(
		"Id:", result["id"],
		"\nName:", result["name"],
		"\nDepartment:", result["department"],
		"\nCity:", address["city"],
		"\nState:", address["state"],
		"\nCountry:", address["country"],
	)
}