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

type ReposAPI interface {
	GetRepos(username string) ([]Repo, error)
}

type Mock struct {}

func(m *Mock) GetRepos(username string) ([]Repo, error){
	return []Repo{
		Repo{
			StargazersCount: 7,
		},
		Repo{
			StargazersCount: 12,
		},
	}, nil
}


type Github struct {
	url ReposAPI
}

func(g *Github) GetRepos(username string) ([]Repo, error){
	res, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/repos", username))
	if err != nil {
		return nil, err
	}

	var repos []Repo
	if err := json.NewDecoder(res.Body).Decode(&repos); err != nil {
		return nil, err
	}

	return repos, nil
}

func GetAverageStarsByUsername(ReposAPI ReposAPI, username string) (float64, error) {
	repos, err := ReposAPI.GetRepos(username)
	if err != nil {
		return 0, err
	}

	// defer res.Body.Close()

	// // log to file
	// f, err := os.OpenFile("github.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Println(err)
	// }
	// defer f.Close()

	// logger := log.New(f, "results", log.LstdFlags)
	// bodyBytes, _ := ioutil.ReadAll(res.Body)
	// logger.Println(string(bodyBytes))
	
	// get dump of response/request
	// body, _ := httputil.DumpResponse(res, true)
	// log.Println(string(body))

	if len(repos) == 0 {
		return 0, nil
	}

	var total int
	for _, v := range repos {
		total += v.StargazersCount
	}

	return float64(total) / float64(len(repos)), nil
}
