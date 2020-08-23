package zounds_test

import (
	"image"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/kulti/zounds"
)

func TestRectXYOrder(t *testing.T) {
	require.Equal(t, zounds.Rect(10, 20, 30, 40), zounds.Rect(30, 20, 10, 40))
	require.Equal(t, zounds.Rect(10, 20, 30, 40), zounds.Rect(10, 40, 30, 20))
	require.Equal(t, zounds.Rect(10, 20, 30, 40), zounds.Rect(30, 40, 10, 20))
}

func TestRectFromImageRect(t *testing.T) {
	require.Equal(t, zounds.Rect(10, 20, 30, 40), zounds.RectFromImageRect(image.Rect(10, 20, 30, 40)))
}
