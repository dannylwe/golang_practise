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
			"designation": "Director",
			"address": {
				"city": "Mumbai",
				"state": "Maharashtra",
				"country": "India"
			}
		},
		{
			"id": 11,
			"name": "Irshad",
			"department": "IT",
			"designation": "Product Manager",
			"address": {
				"city": "Mumbai",
				"state": "Maharashtra",
				"country": "India"
			}
		},
		{
			"id": 12,
			"name": "Pankaj",
			"department": "IT",
			"designation": "Team Lead",
			"address": {
				"city": "Pune",
				"state": "Maharashtra",
				"country": "India"
			}
		}
	]`

	var results []map[string]interface{}

	json.Unmarshal([]byte(array), &results)

	for key, val := range results {
		fmt.Println("\nreading value for key", key)
		address := val["address"].(map[string]interface{})

		fmt.Println(
			"Id:", val["id"],
			"\nName:", val["name"],
			"\nCity:", address["city"],
			"\nState:", address["state"],
			"\nCountry:", address["country"],
		)
	}
}
