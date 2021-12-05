package cmd

import (
	"os"

	"github.com/slack-go/slack"
)

func getClient() *slack.Client {
	return slack.New(os.Getenv("SLACK_OAUTH_TOKEN"))
}
