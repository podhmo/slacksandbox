package slacknotifier

import (
	"context"

	"github.com/podhmo/slacksandbox/slacksandbox/examples/notify/infra/slack"
)

// Notifier :
type Notifier struct {
	Client   slack.Client
	Channels ChannelsConfig
}

// New :
func New(c Config) *Notifier {
	return &Notifier{
		Client:   slack.New(c.Config),
		Channels: c.Channels,
	}
}

// NotifyWhenAccessed :
func (n *Notifier) NotifyWhenAccessed(ctx context.Context) {
	n.Client.PostMessage(ctx, n.Channels.Accessed, "accessed (o_0)")
}
