package slack

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

type dummyPostMessageEvent struct {
	channel string
	message string
}

type dummyPostMessageClient struct {
	Box           []dummyPostMessageEvent
	*actualClient // xxx: hack
}

func (cl *dummyPostMessageClient) PostMessage(ctx context.Context, channel, message string) error {
	cl.Box = append(cl.Box, dummyPostMessageEvent{channel: channel, message: message})
	return nil
}

func TestNotifyMessage(t *testing.T) {
	makeTarget := func(client Client) *Notifier {
		return &Notifier{
			Channels: ChannelsConfig{
				Accessed: "#accessed",
			},
			Client: client,
		}
	}

	t.Run("accessed", func(t *testing.T) {
		ctx := context.Background()
		dummyClient := &dummyPostMessageClient{}
		notifier := makeTarget(dummyClient)

		notifier.NotifyWhenAccessed(ctx)

		assert.Len(t, dummyClient.Box, 1)
		assert.Exactly(t, "#accessed", dummyClient.Box[0].channel)
		assert.Exactly(t, "accessed (o_0)", dummyClient.Box[0].message)
	})
}
