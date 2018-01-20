package slacknotifier

import (
	"github.com/podhmo/slacksandbox/slacksandbox/examples/notify/infra/slack"
)

// Config :
type Config struct {
	Channels ChannelsConfig `json:"channels"`
	slack.Config
}

// ChannelsConfig :
type ChannelsConfig struct {
	Accessed string `json:"accessed"`
}
