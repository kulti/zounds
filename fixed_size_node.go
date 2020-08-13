package zounds

import "image"

// FixedSizeNode is a simple type to store node size.
// It embeds PositionedNode to implement Bounds method.
type FixedSizeNode struct {
	positionedNode
	width, heigth int
}

// NewFixedSizeNode creates a new instance of FixedSizeNode.
func NewFixedSizeNode(r image.Rectangle) FixedSizeNode {
	return FixedSizeNode{
		positionedNode: positionedNode{p: r.Min},
		width:          r.Dx(),
		heigth:         r.Dy(),
	}
}

// Bounds implements StaticNode.Bounds method.
func (n FixedSizeNode) Bounds() image.Rectangle {
	return image.Rectangle{
		Min: n.p,
		Max: image.Point{X: n.p.X + n.width, Y: n.p.Y + n.heigth},
	}
}
