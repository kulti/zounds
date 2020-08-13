package zounds

import "image"

// positionedNode is a simple type to store and update node position.
type positionedNode struct {
	p image.Point
}

// UpdatePosition implements the MovableNode.UpdatePosition method.
func (n *positionedNode) UpdatePosition(p image.Point) {
	n.p = p
}
