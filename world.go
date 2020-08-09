package zounds

import (
	"image"
	"sort"
)

// StaticNode is an interface for static node.
// StaticNode can be only drawn.
// StaticNode cannot be updated by the World.
type StaticNode interface {
	Bounds() image.Rectangle
	Draw()
}

// World represents world base struct.
// It used to create and manipulate all game objects.
type World struct {
	backgroundNode StaticNode
	nodes          []StaticNode
}

// NewWorld creates new World instance.
func NewWorld() *World {
	return &World{}
}

// SetBackground sets static background image.
func (w *World) SetBackground(node StaticNode) {
	w.backgroundNode = node
}

// AddStaticNode adds static node to the World.
func (w *World) AddStaticNode(node StaticNode) {
	nodeBounds := node.Bounds()
	i := sort.Search(len(w.nodes), func(i int) bool {
		return (nodeBounds.Max.Y > w.nodes[i].Bounds().Max.Y) ||
			(nodeBounds.Max.Y == w.nodes[i].Bounds().Max.Y && nodeBounds.Max.X > w.nodes[i].Bounds().Max.X)
	})
	if i == len(w.nodes) {
		w.nodes = append(w.nodes, node)
	} else {
		w.nodes = append(w.nodes, nil)
		copy(w.nodes[i+1:], w.nodes[i:])
		w.nodes[i] = node
	}
}

// Draw draws the World and all visible nodes on screen.
func (w *World) Draw() {
	if w.backgroundNode != nil {
		w.backgroundNode.Draw()
	}

	for _, node := range w.nodes {
		node.Draw()
	}
}
