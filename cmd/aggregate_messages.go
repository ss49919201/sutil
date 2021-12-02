package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/slack-go/slack"
	"github.com/urfave/cli/v2"
)

func AggregateMessages(cliContext *cli.Context) error {
	ctx := cliContext.Context
	client := getClient()

	res, err := client.GetConversationHistoryContext(ctx, &slack.GetConversationHistoryParameters{
		ChannelID: os.Getenv("CHANNEL_ID"),
		Latest:    os.Getenv("LATEST"),
		Oldest:    os.Getenv("OLDEST"),
	})
	if err != nil {
		return err
	}

	// TODO:
	// if res.HasMore {}

	p := newPrinter()
	for i := range res.Messages {
		user, err := client.GetUserInfo(res.Messages[i].User)
		if err != nil {
			return err
		}
		fmt.Fprintln(p.tw, user.Name)
		fmt.Fprintln(p.tw, res.Messages[i].Text)
	}
	return nil
}

type printer struct {
	tw *tabwriter.Writer
}

func newPrinter() *printer {
	return &printer{
		tw: tabwriter.NewWriter(os.Stdout, 0, 4, 0, ' ', 0),
	}
}

func (p *printer) Print() error {
	return p.tw.Flush()
}
