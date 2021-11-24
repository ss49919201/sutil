package main

import "github.com/slack-go/slack"

func (c *Client) GetConversationHistory(parms *slack.GetConversationHistoryParameters) (*slack.GetConversationHistoryResponse, error) {
	return c.slackClient.GetConversationHistory(parms)
}
