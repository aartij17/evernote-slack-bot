package slackbot

import (
	"github.com/nlopes/slack"
	"time"
)

func DailyTicker() {
	ticker := time.NewTicker(4 * time.Second)
	for {
		select {
		case <-ticker.C:
			postMessage()
		}
	}
}

func postMessage() {
	for _, channel := range AllChanData.Channels {
		params := slack.PostMessageParameters{EscapeText: false}
		API.PostMessage(channel, "/topic hi", params)
	}
}