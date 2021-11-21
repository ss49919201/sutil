package main

import (
	"os"

	"github.com/slack-go/slack"
)

func main() {
	_ = slack.New(os.Getenv("SLACK_TOKEN"))
}
