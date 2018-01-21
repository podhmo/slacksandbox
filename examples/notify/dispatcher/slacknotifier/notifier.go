package slacknotifier

import (
	"bytes"
	"context"
	"time"

	"text/template"

	"github.com/podhmo/slacksandbox/slacksandbox/examples/notify/assets"
	"github.com/podhmo/slacksandbox/slacksandbox/examples/notify/infra/slack"
)

// Notifier :
type Notifier struct {
	Client   slack.Client
	Channels ChannelsConfig
}

// New :
func New(c Config) *Notifier {
	return &Notifier{
		Client:   slack.New(c.Config),
		Channels: c.Channels,
	}
}

// NotifyWhenAccessed :
func (n *Notifier) NotifyWhenAccessed(ctx context.Context) {
	var b bytes.Buffer
	data := map[string]interface{}{
		"now": time.Now(),
	}
	tmplAccessed.Execute(&b, data)
	n.Client.PostMessage(ctx, n.Channels.Accessed, b.String())
}

func mustParse(assetName string, name string) *template.Template {
	b := assets.MustAsset(assetName)
	return template.Must(template.New(name).Parse(string(b))) // using unsafe cast?
}

var tmplAccessed *template.Template

func init() {
	tmplAccessed = mustParse("templates/slack/accessed.tmpl", "accessed")
}
