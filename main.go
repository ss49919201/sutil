package main

import (
	"fmt"
	"log"
	"os"

	"github.com/slack-go/slack"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:   "slack-util",
		Usage:  "", // TODO
		Action: countMessage,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func countMessage(ctx *cli.Context) error {
	client := slack.New(os.Getenv("OAUTH_TOKEN"))
	res, err := client.GetConversationHistoryContext(ctx.Context, &slack.GetConversationHistoryParameters{
		ChannelID: os.Getenv("CHANNEL_ID"),
		Latest:    os.Getenv("LATEST"),
		Oldest:    os.Getenv("OLDEST"),
	})
	if err != nil {
		log.Fatal(err)
	}

	// TODO:
	// if res.HasMore {}
	for i := range res.Messages {
		user, err := client.GetUserInfo(res.Messages[i].User)
		if err != nil {
			return err
		}
		fmt.Println(user.Name)
		fmt.Println(res.Messages[i].Text)
	}
	return nil
}
