package testfixture

import (
	"github.com/podhmo/slacksandbox/slacksandbox/examples/notify/dispatcher/mocks"
	"github.com/stretchr/testify/mock"
)

// NewDispatcher :
func NewDispatcher() *mocks.Dispatcher {
	md := &mocks.Dispatcher{}
	md.On("DispatchAccessed", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	return md
}
