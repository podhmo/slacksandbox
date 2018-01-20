package slack

import (
	"context"

	goslack "github.com/lestrrat/go-slack"
	"github.com/pkg/errors"
)

// Client :
type Client interface {
	postMessageClient
	Hello() string
}

// postMessageClient
type postMessageClient interface {
	PostMessage(ctx context.Context, channel, message string) error
}

// actualClient :
type actualClient struct {
	token string
	Debug bool
}

// Hello : xxx
func (c *actualClient) Hello() string {
	return "hello"
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
