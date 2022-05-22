package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3562826309572-3573108994753-OusCuERHAE5m3kHSqu3VSH52")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03GJQDPQ2Y-3545888513543-7f52156ab5c67cd6a9fe2e9a058306371db209b938f974988c0a285e51869f8b")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("My YOB is <year>", &slacker.CommandDefinition{
		Description: "YOB Calculator",
		Example:     "My YOB is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			age := 2022 - yob
			r := fmt.Sprint("age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancle := context.WithCancel(context.Background())
	defer cancle()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
