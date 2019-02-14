package main

import (
	"flag"
	"os"
	"regexp"
	"time"

	"github.com/nlopes/slack"
	"github.com/tenntenn/goplayground"
)

var re = regexp.MustCompile("^hey gopher\n```[A-Za-z]*\n([\\s\\S]*?\n)```")

func run(api *slack.Client) int {
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				text := ev.Text

				ree := re.Copy()
				if !ree.MatchString(text) {
					continue
				}

				var cli goplayground.Client
				r, err := cli.Run(text)
				if err != nil {
					continue
				}

				if r.Errors != "" {
					sendMessage(rtm, r.Errors, ev.Channel)
					continue
				}

				for i := range r.Events {
					time.Sleep(r.Events[i].Delay)
					sendMessage(rtm, r.Events[i].Message, ev.Channel)
				}
			}
		}
	}
}

func main() {
	var (
		token string
	)

	flag.StringVar(&token, "token", "YOUR_API_TOKEN_HERE", "Your Slack Bot API Token")
	flag.Parse()

	api := slack.New(token)
	os.Exit(run(api))
}
