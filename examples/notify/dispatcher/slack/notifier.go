package slack

import "context"

// Notifier :
type Notifier struct {
	Client   Client
	Channels ChannelsConfig
}

// New :
func New(c Config) *Notifier {
	return &Notifier{
		Client:   &actualClient{token: c.Token, Debug: c.Debug},
		Channels: c.Channels,
	}
}

// NotifyWhenAccessed :
func (n *Notifier) NotifyWhenAccessed(ctx context.Context) {
	n.Client.PostMessage(ctx, n.Channels.Accessed, "accessed (o_0)")
}
