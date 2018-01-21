package slack

import "context"

// PostMessage : for test
func PostMessage(postMessage func(ctx context.Context, channel, message string) error) Client {
	return &overridePostMessage{fn: postMessage}
}

type overridePostMessage struct {
	fn func(ctx context.Context, channel, message string) error
	*actualClient
}

func (c *overridePostMessage) PostMessage(ctx context.Context, channel, message string) error {
	return c.fn(ctx, channel, message)
}
