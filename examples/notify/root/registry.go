package root

import (
	"io"
	"os"

	"github.com/podhmo/slacksandbox/slacksandbox/examples/notify/dispatcher"
)

// Registry :
type Registry struct {
	Output     io.Writer
	Dispatcher dispatcher.Dispatcher
}

// NewRegistry :
func NewRegistry(options ...func(*Registry)) *Registry {
	r := &Registry{}
	for _, op := range options {
		op(r)
	}
	if r.Dispatcher == nil {
		// xxx:
		panic("dispatcher is nil")
	}
	if r.Output == nil {
		r.Output = os.Stdout
	}
	return r
}

// WithDispatcher :
func WithDispatcher(d dispatcher.Dispatcher) func(*Registry) {
	return func(r *Registry) {
		r.Dispatcher = d
	}
}

// WithOutput :
func WithOutput(w io.Writer) func(*Registry) {
	return func(r *Registry) {
		r.Output = w
	}
}
