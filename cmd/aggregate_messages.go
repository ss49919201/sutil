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

	parm := &slack.GetConversationHistoryParameters{
		ChannelID: os.Getenv("CHANNEL_ID"),
		Latest:    os.Getenv("LATEST"),
		Oldest:    os.Getenv("OLDEST"),
	}
	res, err := client.GetConversationHistoryContext(ctx, parm)
	if err != nil {
		return err
	}

	// TODO:
	// if res.HasMore {}

	userMsgCntMap := make(map[string]int)
	for i := range res.Messages {
		user, err := client.GetUserInfo(res.Messages[i].User)
		if err != nil {
			return err
		}
		if _, b := userMsgCntMap[user.Name]; b {
			userMsgCntMap[user.Name]++
		} else {
			userMsgCntMap[user.Name] = 0
		}
	}

	p := newPrinter()
	for name, cnt := range userMsgCntMap {
		fmt.Fprint(p.tw, name)
		fmt.Fprintf(p.tw, "\t%d\n", cnt)
	}
	return p.tw.Flush()
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
