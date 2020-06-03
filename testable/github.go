package testable

import (
	"encoding/json"
	"fmt"
	// "log"

	// "io/ioutil"
	// "log"
	"net/http"
	// "net/http/httputil"
	// "os"
)

type Repo struct {
	StargazersCount int `json:"stargazers_count"`
}

type GetReposAPI interface {
	GetRepos(username string) ([]Repo, error)
}

type Mock struct{}

func (m *Mock) GetRepos(username string) ([]Repo, error) {
	return []Repo{
		Repo{
			StargazersCount: 12,
		},
		Repo{
			StargazersCount: 12,
		},
		Repo{
			StargazersCount: 12,
		},
	}, nil
}

type Github struct {}

func (g *Github) GetRepos(username string) ([]Repo, error){
	res, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/repos", username))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	repos := []Repo{}
	if err := json.NewDecoder(res.Body).Decode(&repos); err != nil {
		return nil, err
	}

	return repos, nil
}

func GetAverageStarsByUsername(GetReposAPI GetReposAPI, username string) (float64, error){
	// res, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/repos", username))
	// if err != nil {
	// 	return 0, err
	// }

	// var repos []Repo
	// if err := json.NewDecoder(res.Body).Decode(&repos); err != nil {
	// 	return 0, nil
	// }
	// fmt.Println(res.Body)
	repos, err := GetReposAPI.GetRepos(username)
	if err != nil {
		return 0, err
	}
	if len(repos) == 0 {
		return 0, nil
	}

	var total int
	for _, v := range repos {
		total += v.StargazersCount
	}

	return float64(total) / float64(len(repos)), nil
}
