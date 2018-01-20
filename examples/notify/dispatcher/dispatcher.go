package dispatcher

import (
	"context"

	"github.com/podhmo/slacksandbox/slacksandbox/examples/notify/dispatcher/slack"
)

//go:generate mockery -name Dispatcher

// Dispatcher :
type Dispatcher interface {
	DispatchAccessed(context.Context)
}

type actualDispatcher struct {
	SlackNotifier *slack.Notifier
}

// New :
func New(slackConfig slack.Config) Dispatcher {
	return &actualDispatcher{
		SlackNotifier: slack.New(slackConfig),
	}
}

// DispatchAccessed :
func (d *actualDispatcher) DispatchAccessed(ctx context.Context) {
	d.SlackNotifier.NotifyWhenAccessed(ctx)
}
