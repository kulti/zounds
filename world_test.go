package zounds_test

import (
	"image"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/kulti/zounds"
)

const (
	nodeWidth  = 32
	nodeHeight = 32
)

func TestWorldNodes(t *testing.T) {
	t.Parallel()

	w := zounds.NewWorld()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	bkgNode := NewMockStaticNode(mockCtl)
	node1 := NewMockStaticNode(mockCtl)
	node2 := NewMockStaticNode(mockCtl)

	w.SetBackground(bkgNode)

	node1.EXPECT().Bounds().AnyTimes()
	w.AddStaticNode(node1)

	node2.EXPECT().Bounds().AnyTimes()
	w.AddStaticNode(node2)

	bkgNode.EXPECT().Draw()
	node1.EXPECT().Draw()
	node2.EXPECT().Draw()
	w.Draw()
}

func TestWorldDrawOrder(t *testing.T) {
	t.Parallel()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	node1Rect := image.Rectangle{Max: image.Point{nodeWidth, nodeHeight}}
	node2Rect := node1Rect
	node2Rect.Min.X++
	node2Rect.Max.X++
	node3Rect := node1Rect
	node3Rect.Min.Y++
	node3Rect.Max.Y++

	node1 := NewMockStaticNode(mockCtl)
	node2 := NewMockStaticNode(mockCtl)
	node3 := NewMockStaticNode(mockCtl)

	node1.EXPECT().Bounds().Return(node1Rect).AnyTimes()
	node2.EXPECT().Bounds().Return(node2Rect).AnyTimes()
	node3.EXPECT().Bounds().Return(node3Rect).AnyTimes()

	addOrders := [][]*MockStaticNode{
		{node1, node2, node3},
		{node1, node3, node2},
		{node2, node1, node3},
		{node2, node3, node1},
		{node3, node1, node2},
		{node3, node2, node1},
	}
	for i, nodes := range addOrders {
		nodes := nodes
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			w := zounds.NewWorld()

			bkgNode := NewMockStaticNode(mockCtl)
			w.SetBackground(bkgNode)

			for _, node := range nodes {
				w.AddStaticNode(node)
			}

			bkgDrawCall := bkgNode.EXPECT().Draw()
			node3DrawCall := node3.EXPECT().Draw().After(bkgDrawCall)
			node2DrawCall := node2.EXPECT().Draw().After(node3DrawCall)
			node1.EXPECT().Draw().After(node2DrawCall)

			w.Draw()
		})
	}
}

func TestWorldDrawOrderEqualNodes(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	node1Rect := image.Rectangle{Max: image.Point{nodeWidth, nodeHeight}}
	node2Rect := node1Rect

	node1 := NewMockStaticNode(mockCtl)
	node2 := NewMockStaticNode(mockCtl)

	node1.EXPECT().Bounds().Return(node1Rect).AnyTimes()
	node2.EXPECT().Bounds().Return(node2Rect).AnyTimes()

	for i, nodes := range [][]*MockStaticNode{
		{node1, node2},
		{node2, node1},
	} {
		nodes := nodes
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			w := zounds.NewWorld()
			for _, node := range nodes {
				w.AddStaticNode(node)
			}

			node0DrawCall := nodes[0].EXPECT().Draw()
			nodes[1].EXPECT().Draw().After(node0DrawCall)

			w.Draw()
		})
	}
}
