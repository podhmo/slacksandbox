package actions

import (
	"context"
	"fmt"

	"github.com/podhmo/slacksandbox/slacksandbox/examples/notify/root"
)

// Accessed :
func Accessed(ctx context.Context, registry *root.Registry, skipped bool) error {
	fmt.Fprintln(registry.Output, "before accessed")
	if skipped {
		return nil
	}

	registry.Dispatcher.DispatchAccessed(ctx)
	fmt.Fprintln(registry.Output, "after accessed")
	return nil
}
