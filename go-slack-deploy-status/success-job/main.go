package main

import (
	"fmt"
	"os"
	"github.com/slack-go/slack"
)


func main() {
	//args := os.Args[1:]
	//fmt.Println(args)

	SLACK_BOT_TOKEN := os.Getenv("SLACK_BOT_TOKEN")
	CHANNEL_ID := os.Getenv("CHANNEL_ID")
	APP_URL := os.Getenv("APP_URL")
	COMMIT_HASH := os.Getenv("COMMIT_HASH")

	api := slack.New(SLACK_BOT_TOKEN)
	preText := "*Hello! Your App Platform Deployemnt has finished!*"
	appURL := "*APP URL:* " + APP_URL 
	deployResult := "*" +  SUCCESS + "*"
	commitHash := "*" + COMMIT_HASH + "*"
	jobName := "*" + POST_DEPLOY_RESULTS + "*"

	if deployResult == "*SUCCESS*" { 
		deployResult = deployResult + " :white_check_mark:"
	} else {
		deployResult = deployResult + " :x:"
	}

	dividerSection1 := slack.NewDividerBlock() 
	appPlatformDeployDetails := jobName + " #" + commitHash + " - " + deployResult + "\n" + appURL 
	preTextField := slack.NewTextBlockObject("mrkdwn", preText + "\n\n", false, false) 
	appPlatformDeployDetailsField := slack.NewTextBlockObject("mrkdwn", appPlatformDeployDetails, false, false)

	appPlatformDeployDetailsSection := slack.NewSectionBlock(appPlatformDeployDetailsField, nil, nil)
	preTextSection := slack.NewSectionBlock(preTextField, nil, nil)

	msg := slack.MsgOptionBlocks(
			preTextSection,
			dividerSection1,
			appPlatformDeployDetailsSection,
	)

	_, _, _, err := api.SendMessage(
		CHANNEL_ID,
		msg,
	)

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}