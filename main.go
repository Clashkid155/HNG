package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	h1 := func(w http.ResponseWriter, req *http.Request) {
		currentTime := time.Now()
		utcFormat := currentTime.UTC().Format("2006-01-02T15:04:05Z")
		fmt.Println(currentTime.UTC())
		retInfo := ReturnInfo{
			SlackName:     req.URL.Query().Get("slack_name"),
			CurrentDay:    currentTime.Weekday().String(),
			UtcTime:       utcFormat,
			Track:         req.URL.Query().Get("track"),
			GithubFileUrl: "",
			GithubRepoUrl: "",
			StatusCode:    http.StatusOK,
		}

		//io.WriteString(w, "Hello from a HandleFunc #1!\n")
		info, err := json.Marshal(retInfo)
		if err != nil {
			log.Println(err)
		}

		//fmt.Fprint(w, info)
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
