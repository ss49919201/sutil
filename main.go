package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gemcook/merr"
	"github.com/slack-go/slack"
)

func main() {
	client := slack.New(os.Getenv("OAUTH_TOKEN"))
	res, err := client.GetConversationHistory(&slack.GetConversationHistoryParameters{
		ChannelID: os.Getenv("CHANNEL_ID"),
		Latest:    os.Getenv("LATEST"),
		Oldest:    os.Getenv("OLDEST"),
	})
	if err != nil {
		log.Fatal(err)
	}

	errs := merr.New()
	for i := range res.Messages {
		user, err := client.GetUserInfo(res.Messages[i].User)
		if err != nil {
			errs.Append(err)
		}
		fmt.Println(user.Name)
		fmt.Println(res.Messages[i].Text)
	}
	if errs != nil {
		log.Println(errs)
	}
}
