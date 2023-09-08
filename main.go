package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const (
	fileUrl = "https://github.com/Clashkid155/HNG/blob/main/main.go"
	repoUrl = "https://github.com/Clashkid155/HNG"
)

func main() {
	h1 := func(w http.ResponseWriter, req *http.Request) {
		currentTime := time.Now()
		utcFormat := currentTime.UTC().Format("2006-01-02T15:04:05Z")

		retInfo := ReturnInfo{
			SlackName:     req.URL.Query().Get("slack_name"),
			CurrentDay:    currentTime.Weekday().String(),
			UtcTime:       utcFormat,
			Track:         req.URL.Query().Get("track"),
			GithubFileUrl: fileUrl,
			GithubRepoUrl: repoUrl,
			StatusCode:    http.StatusOK,
		}

		info, err := json.Marshal(retInfo)
		if err != nil {
			log.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(info)
		if err != nil {
			log.Println(err)
		}
	}

	http.HandleFunc("/api", h1)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type ReturnInfo struct {
	SlackName     string `json:"slack_name"`
	CurrentDay    string `json:"current_day"`
	UtcTime       string `json:"utc_time"`
	Track         string `json:"track"`
	GithubFileUrl string `json:"github_file_url"`
	GithubRepoUrl string `json:"github_repo_url"`
	StatusCode    int    `json:"status_code"`
}
