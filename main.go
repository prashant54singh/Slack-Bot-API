package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

// SLACK external package to help me talk to external library, someone has already made the network request with wrapper around it
func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3999134615701-4004577158244-IH8YTjJoE9qMVfjb8ayiO9zB")
	os.Setenv("CHANNEL_ID", "C03VD3YJAS1")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"cntask.pdf"}
	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URL)
	}
}
