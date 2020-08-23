package zounds

// positionedNode is a simple type to store and update node position.
type positionedNode struct {
	p Point
}

// UpdatePosition implements the MovableNode.UpdatePosition method.
func (n *positionedNode) UpdatePosition(p Point) {
	n.p = p
}
