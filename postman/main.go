package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/lestrrat/go-slack"
	"github.com/pkg/errors"
)

func run(token string, w io.Writer) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cl := slack.New(token, slack.WithDebug(true))

	// check if we are connected
	authres, err := cl.Auth().Test().Do(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to test authentication")
	}
	fmt.Fprintf(w, "%#v\n", authres)

	// simplest possible message
	chatres, err := cl.Chat().PostMessage("#api-sandbox").
		Text("Hello, World!").
		Do(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to post messsage")
	}
	fmt.Fprintf(w, "%#v\n", chatres)
	return nil
}

func main() {
	token := os.Getenv("SLACK_TOKEN")
	if err := run(token, os.Stdout); err != nil {
		log.Fatal(err)
	}
}
