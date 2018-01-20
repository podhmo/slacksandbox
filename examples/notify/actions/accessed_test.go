package actions

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/podhmo/slacksandbox/slacksandbox/examples/notify/dispatcher/mocks"
	"github.com/podhmo/slacksandbox/slacksandbox/examples/notify/root"
	"github.com/podhmo/slacksandbox/slacksandbox/examples/notify/testfixture"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccessedDispatcher(t *testing.T) {
	type c struct {
		msg                   string
		skipped               bool
		expectedNumberofCalls int
	}
	candidates := []c{
		{
			msg:                   "not skipped, called",
			skipped:               false,
			expectedNumberofCalls: 1,
		},
		{
			msg:                   "skipped, not called",
			skipped:               true,
			expectedNumberofCalls: 0,
		},
	}

	for _, c := range candidates {
		c := c
		t.Run(c.msg, func(t *testing.T) {
			md := &mocks.Dispatcher{}
			ctx := context.Background()
			md.On("DispatchAccessed", ctx).Return(nil)

			registry := testfixture.NewRegistry(root.WithDispatcher(md))

			err := Accessed(context.Background(), registry, c.skipped)
			require.NoError(t, err)
			md.AssertNumberOfCalls(t, "DispatchAccessed", c.expectedNumberofCalls)
		})
	}
}

func TestAccessedOutput(t *testing.T) {
	t.Run("not skipped, full response", func(t *testing.T) {
		var b bytes.Buffer
		registry := testfixture.NewRegistry(root.WithOutput(&b))

		skipped := false
		err := Accessed(context.Background(), registry, skipped)
		require.NoError(t, err)

		output := b.String()
		assert.True(t, strings.Contains(output, "before accessed"))
		assert.True(t, strings.Contains(output, "after accessed"))
	})

	t.Run("skipped, only before", func(t *testing.T) {
		var b bytes.Buffer
		registry := testfixture.NewRegistry(root.WithOutput(&b))

		skipped := true
		err := Accessed(context.Background(), registry, skipped)
		require.NoError(t, err)

		output := b.String()
		assert.True(t, strings.Contains(output, "before accessed"))
		assert.False(t, strings.Contains(output, "after accessed"))
	})
}
