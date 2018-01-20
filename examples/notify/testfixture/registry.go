package testfixture

import (
	"github.com/podhmo/slacksandbox/slacksandbox/examples/notify/root"
)

var null NullWriter

// NewRegistry :
func NewRegistry(options ...func(*root.Registry)) *root.Registry {
	options = append(options, func(r *root.Registry) {
		if r.Output == nil {
			r.Output = null
		}
		if r.Dispatcher == nil {
			r.Dispatcher = NewDispatcher()
		}
	})
	return root.NewRegistryWithOptions(options...)
}
