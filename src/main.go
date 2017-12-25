package main

import (
	"flag"
	"github.com/nlopes/slack"
	"slackbot"
)

func main() {

	var keyFile string
	flag.StringVar(&keyFile, "key", "", "Secrets file containing API Token")
	flag.Parse()

	apiToken := slackbot.GetAPIToken(keyFile)

	// read the key from a secrets file
	slackbot.API = slack.New(apiToken)

	slackbot.BuildChannelMetaData()
	slackbot.DailyTicker()
}
