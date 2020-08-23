package zebiten

import (
	"github.com/hajimehoshi/ebiten"

	"github.com/kulti/zounds"
)

// ImageNode contains static image and draws it.
// It implements zounds.StaticNode interface.
type ImageNode struct {
	zounds.FixedSizeNode
	screen   *screen
	img      *ebiten.Image
	drawOpts *ebiten.DrawImageOptions
}

// NewImageNode creates a new ImageNode instance.
func (g *Game) NewImageNode(img *ebiten.Image, pos zounds.Point) *ImageNode {
	drawOpts := &ebiten.DrawImageOptions{}
	drawOpts.GeoM.Translate(pos.X, pos.Y)

	return &ImageNode{
		screen:        g.screen,
		img:           img,
		FixedSizeNode: zounds.NewFixedSizeNode(zounds.RectFromImageRect(img.Bounds()).Add(pos)),
		drawOpts:      drawOpts,
	}
}

// Draw draws image node.
func (n *ImageNode) Draw() {
	n.screen.DrawImage(n.img, n.drawOpts)
}
