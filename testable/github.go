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

func GetAverageStarsByUsername(username string) (float64, error) {
	res, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/repos", username))
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
	
	var repos []Repo
	if err := json.NewDecoder(res.Body).Decode(&repos); err != nil {
		return 0, nil
	}

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
