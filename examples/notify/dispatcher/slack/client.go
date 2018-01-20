package slack

import (
	"context"

	goslack "github.com/lestrrat/go-slack"
	"github.com/pkg/errors"
)

// todo: debug option

// Client :
type Client interface {
	PostMessage(ctx context.Context, channel, message string) error
}

// actualClient :
type actualClient struct {
	token string
	Debug bool
}

func (c *actualClient) PostMessage(ctx context.Context, channel, message string) error {
	// todo: reuse
	cl := goslack.New(c.token, goslack.WithDebug(c.Debug))
	_, err := cl.Auth().Test().Do(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to test authentication")
	}
	_, err = cl.Chat().
		PostMessage(channel).
		Text(message).
		Do(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to post messsage")
	}
	return nil
}
