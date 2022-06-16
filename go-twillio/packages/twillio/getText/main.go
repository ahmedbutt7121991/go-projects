package main

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"github.com/twillio/twillio-go"
	openapi "github.com/twillio/twillio-go/rest/api/v2010"
)

var (
	accountSid string
	authToken string
	fromPhone string
	toPhone string
	client *twillio.RestClient
)

func SendMessage(msg string) {
	params := openapi.CreateMessageparams{}
	params.SetTo(toPhone)
	params.SetFrom(fromPhone)
	params.SetBody(msg)

	response, err := client.ApiV2010.CreateMessage(&params)
	if err != nil {
		fmt.Printf("error creatign and sending message: %s\n", err.Error())
		return
	}
	fmt.Printf("Message SID: %s\n", *response.Sid)
}

func init() {
	err := godotenv.Load(".env")
	if err != nil{
		fmt.Printf("error loading .env: %s\n", err.Error())
		os.Exit(1)
	}
	accountSid = os.Getenv("ACCOUNT_SID")
	authToken = os.Getenv("AUTH_TOKEN")
	fromPhone = os.Getenv("FROM_PHONE")
	toPhone = os.Getenv("TO_PHONE")

	client = twillio.NewRestClientWithParams(twilio.RestClientParams{
		Username: accountSid,
		Password: authToken,
	})
}

func main() {
	msg := fmt.Sprintf(os.Getenv("MSG"), "AHMED")
	SendMessage(msg)
}