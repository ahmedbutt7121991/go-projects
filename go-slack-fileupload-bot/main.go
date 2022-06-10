package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {

	//os.Setenv("SLACK_BOT_TOKEN", "")
	//os.Setenv("CHANNEL_ID", "")
	
	SLACK_BOT_TOKEN := os.Getenv("SLACK_BOT_TOKEN")
	CHANNEL_ID := os.Getenv("CHANNEL_ID")
	
	
	api := slack.New(SLACK_BOT_TOKEN)
	channelArr := []string{CHANNEL_ID}
	fileArr := []string{"docker-cheat-sheet.pdf"}

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
