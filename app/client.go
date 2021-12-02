package app

import (
	"os"

	"github.com/slack-go/slack"
)

func getClient() *slack.Client {
	return slack.New(os.Getenv("OAUTH_TOKEN"))
}
