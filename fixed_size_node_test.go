package zounds_test

import (
	"image"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/kulti/zounds"
)

func TestFixSizedNode(t *testing.T) {
	t.Parallel()

	randomRect := image.Rect(rand.Int(), rand.Int(), rand.Int(), rand.Int()) //nolint:gosec
	node := zounds.NewFixedSizeNode(randomRect)
	require.Equal(t, randomRect, node.Bounds())
}
