package slacknotifier

import (
	"context"
	"testing"

	"github.com/podhmo/slacksandbox/slacksandbox/examples/notify/infra/slack"
	"github.com/stretchr/testify/assert"
)

func TestNotifyMessage(t *testing.T) {
	type dummyPostEvent struct {
		channel string
		message string
	}

	makeTarget := func(box *[]dummyPostEvent) *Notifier {
		return &Notifier{
			Channels: ChannelsConfig{
				Accessed: "#accessed",
			},
			Client: slack.PostMessage(func(ctx context.Context, channel, message string) error {
				*box = append(*box, dummyPostEvent{channel: channel, message: message})
				return nil
			}),
		}
	}

	t.Run("accessed", func(t *testing.T) {
		ctx := context.Background()
		box := []dummyPostEvent{}
		notifier := makeTarget(&box)

		notifier.NotifyWhenAccessed(ctx)

		assert.Len(t, box, 1)
		assert.Exactly(t, "#accessed", box[0].channel)
		assert.Exactly(t, "accessed (o_0)", box[0].message)
	})
}
