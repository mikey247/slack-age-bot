package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent)  {
	for event := range analyticsChannel{
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
	}
}

func main()  {
	os.Setenv("SLACK_BOT_TOKEN","xoxb-4303658227253-4320829293713-NWzl2pvAmgnuwNo1YD6MYDzG")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A0494FM0GCC-4305133030773-633dace677f4a19617fd54a71facf8e3b695dd33b2b9519d8c247c4ae764bab1")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Examples:[]string{"my yob is 2020"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			age := time.Now().Year() - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	ctx,cancel := context.WithCancel(context.Background())
	defer cancel()

	err:= bot.Listen(ctx)
	if err !=nil{
		log.Fatal(err)
	}
}