package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/podhmo/slacksandbox/slacksandbox/examples/notify/dispatcher"
	"github.com/podhmo/slacksandbox/slacksandbox/examples/notify/dispatcher/slack"
)

type config struct {
	Slack slack.Config `json:"slack"`
}

type app struct {
	Dispatcher dispatcher.Dispatcher
}

func makeApp(confpath string) (*app, error) {
	f, err := os.Open(confpath)
	if err != nil {
		return nil, errors.Wrap(err, "load config")
	}

	var c config
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&c); err != nil {
		return nil, errors.Wrap(err, "decode config")
	}

	return &app{
		Dispatcher: dispatcher.New(c.Slack),
	}, nil
}

func run(confpath string) error {
	app, err := makeApp(confpath)
	if err != nil {
		return err
	}

	ctx := context.Background()
	app.Dispatcher.DispatchAccessed(ctx)

	return nil
}

func main() {
	confpath := os.Getenv("CONF")
	if err := run(confpath); err != nil {
		log.Fatalf("error: #%v", err)
	}
}
