package main

import (
	"os"

	"github.com/slack-go/slack"
)

type Client struct {
	slackClient *slack.Client
}

func NewClient() *Client {
	return &Client{slack.New(os.Getenv("OAUTH_TOKEN"))}
}
