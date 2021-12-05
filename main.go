package main

import (
	"os"
	"strconv"
	"time"

	"github.com/s-beats/sutil/cmd"
	"github.com/s-beats/sutil/logger"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "sutil"
	app.Usage = "A collection of useful commands for working with Slack."

	app.Commands = []*cli.Command{
		{
			Name:   "aggregate-messages",
			Action: cmd.AggregateMessages,
			Flags: []cli.Flag{
				&cli.Int64Flag{
					Name:        "from",
					Aliases:     []string{"f"},
					DefaultText: "0",
				},
				&cli.Int64Flag{
					Name:        "to",
					Aliases:     []string{"t"},
					DefaultText: strconv.Itoa(int(time.Now().Unix())),
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Error(err)
	}
}
