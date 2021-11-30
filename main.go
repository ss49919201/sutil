package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gemcook/merr"
	"github.com/slack-go/slack"
)

func main() {
	ctx := context.Background()

	client := slack.New(os.Getenv("OAUTH_TOKEN"))
	res, err := client.GetConversationHistoryContext(ctx, &slack.GetConversationHistoryParameters{
		ChannelID: os.Getenv("CHANNEL_ID"),
		Latest:    os.Getenv("LATEST"),
		Oldest:    os.Getenv("OLDEST"),
	})
	if err != nil {
		log.Fatal(err)
	}

	// TODO:
	// if res.HasMore {}

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
