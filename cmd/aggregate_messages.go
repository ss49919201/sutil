package cmd

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/rs/zerolog/log"
	"github.com/slack-go/slack"
	"github.com/urfave/cli/v2"
)

func AggregateMessages(cliContext *cli.Context) error {
	ctx := cliContext.Context
	from := cliContext.Int64("from")
	to := cliContext.Int64("to")

	log.Info().Msg(fmt.Sprintf("aggregate %d ~ %d", from, to))

	client := getClient()
	parm := &slack.GetConversationHistoryParameters{
		ChannelID: os.Getenv("CHANNEL_ID"),
		Latest:    strconv.Itoa(int(to)),
		Oldest:    strconv.Itoa(int(from)),
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

	// set column title
	fmt.Fprint(p.tw, "USER NAME")
	fmt.Fprint(p.tw, "\tMESSAGES COUNT\n")

	for name, cnt := range userMsgCntMap {
		fmt.Fprint(p.tw, name)
		fmt.Fprintf(p.tw, "\t%d\n", cnt)
	}
	return p.Print()
}

type printer struct {
	tw *tabwriter.Writer
}

func newPrinter() *printer {
	return &printer{
		tw: tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0),
	}
}

func (p *printer) Print() error {
	return p.tw.Flush()
}
