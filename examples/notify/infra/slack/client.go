package slack

import (
	"context"

	goslack "github.com/lestrrat/go-slack"
	"github.com/pkg/errors"
)

// Client :
type Client interface {
	postMessage
	Hello() string
}

// New :
func New(c Config) Client {
	return &ActualClient{token: c.Token, Debug: c.Debug}
}

// postMessage :
type postMessage interface {
	PostMessage(ctx context.Context, channel, message string) error
}

// ActualClient :
type ActualClient struct {
	token string
	Debug bool
}

// Hello : xxx
func (c *ActualClient) Hello() string {
	return "hello"
}

// PostMessage :
func (c *ActualClient) PostMessage(ctx context.Context, channel, message string) error {
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
