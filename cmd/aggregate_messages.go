package cmd

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/slack-go/slack"
	"github.com/urfave/cli/v2"
)

type userMessagesCount struct {
	userName string
	count    int
}

type userMessagesCountList struct {
	data []*userMessagesCount
}

func newUserMessagesCountList() *userMessagesCountList {
	return &userMessagesCountList{
		data: make([]*userMessagesCount, 0),
	}
}

func (u *userMessagesCountList) existUser(userName string) bool {
	for _, v := range u.data {
		if v.userName == userName {
			return true
		}
	}
	return false
}

func (u *userMessagesCountList) increment(userName string) *userMessagesCountList {
	for _, v := range u.data {
		if v.userName == userName {
			v.count++
			return u
		}
	}
	return u
}

func (u *userMessagesCountList) appendUser(userName string) *userMessagesCountList {
	u.data = append(u.data, &userMessagesCount{
		userName: userName,
	})
	return u
}

func getConversationHistory(ctx context.Context, chanID, latest, oldest string) (*slack.GetConversationHistoryResponse, error) {
	client := ctx.Value(slackClientContextKey).(*slack.Client)
	parm := &slack.GetConversationHistoryParameters{
		ChannelID: chanID,
		Latest:    latest,
		Oldest:    oldest,
	}
	res, err := client.GetConversationHistoryContext(ctx, parm)
	if err != nil {
		return nil, err
	}
	// TODO:
	// if res.HasMore {...}
	return res, nil
}

func getUser(ctx context.Context, user string) (*slack.User, error) {
	client := ctx.Value(slackClientContextKey).(*slack.Client)
	return client.GetUserInfoContext(ctx, user)
}

func print(list *userMessagesCountList) error {
	p := newPrinter()
	p.initTable()
	for _, v := range list.data {
		fmt.Fprint(p.tw, v.userName)
		fmt.Fprintf(p.tw, "\t%d\n", v.count)
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

// set column title
func (p *printer) initTable() {
	fmt.Fprint(p.tw, "USER NAME")
	fmt.Fprint(p.tw, "\tMESSAGES COUNT\n")
}

// Print print bufferd data
func (p *printer) Print() error {
	return p.tw.Flush()
}

// AggregateMessages aggreagate messages of each user
func AggregateMessages(cliContext *cli.Context) error {
	ctx := cliContext.Context
	ctx = context.WithValue(ctx, slackClientContextKey, getClient())

	chanID := cliContext.String("chanid")
	from := cliContext.Int64("from")
	to := cliContext.Int64("to")

	res, err := getConversationHistory(ctx, chanID, strconv.Itoa(int(to)), strconv.Itoa(int(from)))
	if err != nil {
		return err
	}

	userMsgCountSet := newUserMessagesCountList()
	for i := range res.Messages {
		user, err := getUser(ctx, res.Messages[i].User)
		if err != nil {
			return err
		}
		if userMsgCountSet.existUser(user.Name) {
			userMsgCountSet = userMsgCountSet.increment(user.Name)
		} else {
			userMsgCountSet = userMsgCountSet.appendUser(user.Name)
		}
	}

	return print(userMsgCountSet)
}
