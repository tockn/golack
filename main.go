package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/nlopes/slack"
	"github.com/tenntenn/goplayground"
)

var(
	re *regexp.Regexp
)

func init() {
	re = regexp.MustCompile("^hey gopher\n```[A-Za-z]*\n([\\s\\S]*?\n)```")
}

func main() {

	token := os.Getenv("TOKEN")

	if token == "" {
		os.Exit(0)
	}

	api := slack.New(token)
	os.Exit(run(api))
}

func run(api *slack.Client) int {
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:

				ree := re.Copy()
				if !ree.MatchString(ev.Text) {
					continue
				}

				src := retrieveSourceCode(ev.Text)

				var cli goplayground.Client
				r, err := cli.Run(src)
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

func retrieveSourceCode(s string) string{
	s = strings.Replace(s, "hey gopher", "", 1)
	s = strings.Replace(s, "```", "", 2)
	return s
}

func sendMessage(rtm *slack.RTM, text string, channelID string, options ...slack.RTMsgOption) {
	s := fmt.Sprintf("```%s```", text)
	rtm.SendMessage(rtm.NewOutgoingMessage(s, channelID, options...))
}

