package zounds

// FixedSizeNode is a simple type to store node size.
// It embeds PositionedNode to implement Bounds method.
type FixedSizeNode struct {
	positionedNode
	width, heigth float64
}

// NewFixedSizeNode creates a new instance of FixedSizeNode.
func NewFixedSizeNode(r Rectangle) FixedSizeNode {
	return FixedSizeNode{
		positionedNode: positionedNode{p: r.Min},
		width:          r.Dx(),
		heigth:         r.Dy(),
	}
}

// Bounds implements StaticNode.Bounds method.
func (n FixedSizeNode) Bounds() Rectangle {
	return Rectangle{
		Min: n.p,
		Max: Point{X: n.p.X + n.width, Y: n.p.Y + n.heigth},
	}
}
