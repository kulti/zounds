package zounds_test

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/kulti/zounds"
)

func TestFixSizedNode(t *testing.T) {
	t.Parallel()

	randomRect := zounds.Rect(rand.Float64(), rand.Float64(), rand.Float64(), rand.Float64()) //nolint:gosec
	node := zounds.NewFixedSizeNode(randomRect)
	require.Equal(t, randomRect, node.Bounds())
}
