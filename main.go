package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/nlopes/slack"
)


func main() {
	apiToken := os.Getenv("SLACK_API_TOKEN")
	if apiToken == "" {
		panic("$SLACK_API_TOKEN is not set")
	}
	api := slack.New(apiToken)

	emojis, err := api.GetEmoji()
	if err != nil {
		panic(err)
	}
	if err := os.MkdirAll("emoji", 0777); err != nil {
		panic(err)
	}

	for name, emojiUrl := range emojis {
		if strings.Index(emojiUrl, "alias") != -1 {
			continue
		}
		response, err := http.Get(emojiUrl)
		if err != nil {
			panic(err)
		}

		file, err := os.Create(filepath.Join("emoji", name + ".png"))
		if err != nil {
			panic(err)
		}
		io.Copy(file, response.Body)
	}
}
