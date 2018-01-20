package dispatcher

import (
	"context"

	"github.com/podhmo/slacksandbox/slacksandbox/examples/notify/dispatcher/slacknotifier"
)

type SlackConfig = slacknotifier.Config

// Dispatcher :
type Dispatcher interface {
	DispatchAccessed(context.Context)
}

type actualDispatcher struct {
	SlackNotifier *slacknotifier.Notifier
}

// New :
func New(slackConfig SlackConfig) Dispatcher {
	return &actualDispatcher{
		SlackNotifier: slacknotifier.New(slackConfig),
	}
}

// DispatchAccessed :
func (d *actualDispatcher) DispatchAccessed(ctx context.Context) {
	d.SlackNotifier.NotifyWhenAccessed(ctx)
}
