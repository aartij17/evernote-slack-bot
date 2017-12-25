package slackbot

import (
	"fmt"
	"github.com/nlopes/slack"
)

type channelData struct {
	Channels []string
}

var API *slack.Client
var AllChanData *channelData

func createChannelData() *channelData {
	return &channelData{}
}

func init() {
	AllChanData = createChannelData()
}

// BuildChannelMetaData gets all the channels in the workspace and saves their IDs.
// These IDs are further needed when we want to post a message to a channel.
func BuildChannelMetaData() {
	channels, err := API.GetChannels(false)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, channel := range channels {
		//fmt.Println(channel.ID)
		AllChanData.Channels = append(AllChanData.Channels, channel.ID)
	}
}
