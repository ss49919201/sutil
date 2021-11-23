package main

import (
	"github.com/slack-go/slack"
)

func (c *Client) PostMessage(channelID, text string, escape bool) error {
	_, _, err := c.slackClient.PostMessage(channelID, slack.MsgOptionText(text, escape))
	if err != nil {
		return err
	}
	return nil
}
