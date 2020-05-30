package testable

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Repo struct {
	StargazersCount int `json:"stargazers_count"`
}

func GetAverageStarsByUsername(username string) (float64, error){
	res, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/repos", username))
	if err != nil {
		return 0, err
	}

	var repos []Repo
	if err := json.NewDecoder(res.Body).Decode(&repos); err != nil {
		return 0, nil
	}
	fmt.Println(res.Body)

	if len(repos) == 0 {
		return 0, nil
	}

	var total int
	for _, v := range repos {
		total += v.StargazersCount
	}

	return float64(total) / float64(len(repos)), nil
}