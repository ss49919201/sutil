package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/s-beats/sutil/cmd"
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
				&cli.StringFlag{
					Name:     "chanid",
					Required: true,
				},
				&cli.Int64Flag{
					Name:        "from",
					DefaultText: "0",
				},
				&cli.Int64Flag{
					Name:        "to",
					DefaultText: strconv.Itoa(int(time.Now().Unix())),
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
