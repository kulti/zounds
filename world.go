package zounds

import (
	"sort"
	"time"

	"github.com/cabify/timex"
)

// StaticNode is an interface for static node.
// StaticNode can be only drawn.
// StaticNode cannot be updated by the World.
type StaticNode interface {
	Bounds() Rectangle
	Draw()
}

// DynamicNode is an interface for dynamic node.
// It behaves like StaticNode but her state can be updated.
// E.g. node with animation can update draw state.
// The delta value in Update method is a time duration since last Update.
type DynamicNode interface {
	StaticNode
	Update(delta time.Duration)
}

// MovableNode is an interface for movable node.
// It should update her state in Update method from DynamicNode interface
// and returns velocity vector from Velocity method.
// The world make desicion how to update node position
// and pass new Min.X, Min.Y by UpdatePosition.
type MovableNode interface {
	DynamicNode
	Velocity() Point
	UpdatePosition(Point)
}

// World represents world base struct.
// It used to create and manipulate all game objects.
type World struct {
	backgroundNode StaticNode
	nodes          []StaticNode
	dynamicNodes   []DynamicNode
	movableNodes   []MovableNode
	lastUpdateTime time.Time
}

// NewWorld creates new World instance.
func NewWorld() *World {
	return &World{}
}

// SetBackground sets static background image.
func (w *World) SetBackground(node StaticNode) {
	w.backgroundNode = node
}

// AddStaticNode adds a static node to the World.
func (w *World) AddStaticNode(node StaticNode) {
	nodeBounds := node.Bounds()
	i := w.searchInsertNodePosition(nodeBounds)
	if i == len(w.nodes) {
		w.nodes = append(w.nodes, node)
	} else {
		w.nodes = append(w.nodes, nil)
		copy(w.nodes[i+1:], w.nodes[i:])
		w.nodes[i] = node
	}
}

// AddDynamicNode adds a dynamic node to the World.
func (w *World) AddDynamicNode(node DynamicNode) {
	w.AddStaticNode(node)
	w.dynamicNodes = append(w.dynamicNodes, node)
}

// AddMovableNode adds a movable node to the World.
func (w *World) AddMovableNode(node MovableNode) {
	w.AddDynamicNode(node)
	w.movableNodes = append(w.movableNodes, node)
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

// Update updates the World and all visible nodes on screen.
func (w *World) Update() {
	var delta time.Duration
	if !w.lastUpdateTime.IsZero() {
		delta = timex.Since(w.lastUpdateTime)
	}
	w.lastUpdateTime = timex.Now()

	for _, node := range w.dynamicNodes {
		node.Update(delta)
	}

	for _, node := range w.movableNodes {
		v := node.Velocity()
		if v.X == 0 && v.Y == 0 {
			continue
		}

		oldIndex := w.searchInsertNodePosition(node.Bounds()) - 1
		newBounds := node.Bounds().Add(v)
		newIndex := w.searchInsertNodePosition(newBounds)
		node.UpdatePosition(node.Bounds().Min.Add(v))
		if oldIndex == newIndex || oldIndex+1 == newIndex {
			continue
		}

		if oldIndex < newIndex {
			copy(w.nodes[oldIndex:newIndex-1], w.nodes[oldIndex+1:newIndex])
			w.nodes[newIndex-1] = node
		} else {
			copy(w.nodes[newIndex+1:oldIndex+1], w.nodes[newIndex:oldIndex])
			w.nodes[newIndex] = node
		}
	}
}

func (w *World) searchInsertNodePosition(nodeBounds Rectangle) int {
	return sort.Search(len(w.nodes), func(i int) bool {
		return (nodeBounds.Max.Y < w.nodes[i].Bounds().Max.Y) ||
			(nodeBounds.Max.Y == w.nodes[i].Bounds().Max.Y && nodeBounds.Max.X < w.nodes[i].Bounds().Max.X)
	})
}
