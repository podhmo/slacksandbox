package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/podhmo/slacksandbox/slacksandbox/examples/notify/actions"
	"github.com/podhmo/slacksandbox/slacksandbox/examples/notify/dispatcher"
	"github.com/podhmo/slacksandbox/slacksandbox/examples/notify/dispatcher/slack"
	"github.com/podhmo/slacksandbox/slacksandbox/examples/notify/root"
)

type config struct {
	Slack slack.Config `json:"slack"` // todo: dispatcher.SlackConfig (type alias)
}

type app struct {
	Registry *root.Registry
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
		Registry: root.NewRegistry(dispatcher.New(c.Slack)),
	}, nil
}

func run(confpath string) error {
	app, err := makeApp(confpath)
	if err != nil {
		return err
	}

	ctx := context.Background()
	skipped := false
	if err := actions.Accessed(ctx, app.Registry, skipped); err != nil {
		return err
	}
	return nil
}

func main() {
	confpath := os.Getenv("CONF")
	if err := run(confpath); err != nil {
		log.Fatalf("error: #%v", err)
	}
}
