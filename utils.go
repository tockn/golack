package main

import (
	"fmt"
	"github.com/nlopes/slack"
)

func sendMessage(rtm *slack.RTM, text string, channelID string, options ...slack.RTMsgOption) {
	s := fmt.Sprintf("```%s```", text)
	rtm.SendMessage(rtm.NewOutgoingMessage(s, channelID, options...))
}
