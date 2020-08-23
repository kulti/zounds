package zounds_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/cabify/timex/timextest"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	"github.com/kulti/zounds"
)

const (
	nodeWidth  = 32
	nodeHeight = 32
)

type WorldSuite struct {
	suite.Suite
	mockCtl *gomock.Controller
	world   *zounds.World
}

func (s *WorldSuite) SetupTest() {
	s.world = zounds.NewWorld()
	s.mockCtl = gomock.NewController(s.T())
}

func (s *WorldSuite) TearDownTest() {
	s.mockCtl.Finish()
}

func (s *WorldSuite) TestWorldDrawOrder() {
	node1Rect := zounds.Rectangle{Max: zounds.Point{nodeWidth, nodeHeight}}
	node2Rect := node1Rect
	node2Rect.Min.X++
	node2Rect.Max.X++
	node3Rect := node1Rect
	node3Rect.Min.Y++
	node3Rect.Max.Y++

	node1 := NewMockStaticNode(s.mockCtl)
	node2 := NewMockStaticNode(s.mockCtl)
	node3 := NewMockStaticNode(s.mockCtl)

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
		s.Run(strconv.Itoa(i), func() {
			w := zounds.NewWorld()

			bkgNode := NewMockStaticNode(s.mockCtl)
			w.SetBackground(bkgNode)

			for _, node := range nodes {
				w.AddStaticNode(node)
			}

			checkDrawOrder(w, bkgNode.EXPECT(), node1.EXPECT(), node2.EXPECT(), node3.EXPECT())
		})
	}
}

func (s *WorldSuite) TestWorldDrawOrderEqualNodes() {
	node1Rect := zounds.Rectangle{Max: zounds.Point{nodeWidth, nodeHeight}}
	node2Rect := node1Rect

	node1 := NewMockStaticNode(s.mockCtl)
	node2 := NewMockStaticNode(s.mockCtl)

	node1.EXPECT().Bounds().Return(node1Rect).AnyTimes()
	node2.EXPECT().Bounds().Return(node2Rect).AnyTimes()

	for i, nodes := range [][]*MockStaticNode{
		{node1, node2},
		{node2, node1},
	} {
		nodes := nodes
		s.Run(strconv.Itoa(i), func() {
			w := zounds.NewWorld()
			for _, node := range nodes {
				w.AddStaticNode(node)
			}

			checkDrawOrder(w, nodes[0].EXPECT(), nodes[1].EXPECT())
		})
	}
}

func (s *WorldSuite) TestDynamicNodeUpdateDelta() {
	now := time.Now()
	timextest.Mocked(now, func(mocked *timextest.TestImplementation) {
		node := NewMockDynamicNode(s.mockCtl)

		node.EXPECT().Bounds().AnyTimes()
		s.world.AddDynamicNode(node)

		mocked.SetNow(now.Add(time.Second))
		node.EXPECT().Update(time.Duration(0))
		s.world.Update()

		mocked.SetNow(now.Add(2 * time.Second))
		node.EXPECT().Update(time.Second)
		s.world.Update()
	})
}

func (s *WorldSuite) TestWorldMovableNodeVelocity() {
	initRect := zounds.Rect(nodeWidth, nodeHeight, 2*nodeWidth, 2*nodeHeight)
	node := NewMockMovableNode(s.mockCtl)

	node.EXPECT().Bounds().Return(initRect).AnyTimes()
	s.world.AddMovableNode(node)

	checkVelocity := func(velocity zounds.Point) {
		expectedRect := initRect.Add(velocity)
		node.EXPECT().Update(gomock.Any())
		node.EXPECT().Velocity().Return(velocity)
		node.EXPECT().UpdatePosition(expectedRect.Min)
		s.world.Update()
	}

	s.Run("vx", func() {
		checkVelocity(zounds.Point{3, 0})
	})

	s.Run("vy", func() {
		checkVelocity(zounds.Point{0, 2})
	})

	s.Run("diag", func() {
		checkVelocity(zounds.Point{-1, -5})
	})

	s.Run("zero", func() {
		node.EXPECT().Update(gomock.Any())
		node.EXPECT().Velocity().Return(zounds.Point{})
		s.world.Update()
	})
}

type MovableNodeStub struct {
	*MockDynamicNode
	zounds.FixedSizeNode
	velocity zounds.Point
}

func (n MovableNodeStub) Bounds() zounds.Rectangle {
	return n.FixedSizeNode.Bounds()
}

func (n MovableNodeStub) Velocity() zounds.Point {
	return n.velocity
}

func (s *WorldSuite) TestWorldMoveNodeChangesDrawOrder() {
	var movableNode *MovableNodeStub
	staticNodes := make([]*MockStaticNode, 0, 4)
	startRect := zounds.Rect(nodeWidth, nodeHeight, 2*nodeWidth, 2*nodeHeight)

	for i := 0; i < 5; i++ {
		rect := startRect.Add(zounds.Point{0, float64(i) * 5})
		if i == 2 {
			movableNode = &MovableNodeStub{
				MockDynamicNode: NewMockDynamicNode(s.mockCtl),
				FixedSizeNode:   zounds.NewFixedSizeNode(rect),
				velocity:        zounds.Point{0, -3},
			}
			movableNode.EXPECT().Bounds().Return(rect).AnyTimes()
			movableNode.EXPECT().Update(gomock.Any()).AnyTimes()
			s.world.AddMovableNode(movableNode)
		} else {
			staticNode := NewMockStaticNode(s.mockCtl)
			staticNode.EXPECT().Bounds().Return(rect).AnyTimes()
			s.world.AddStaticNode(staticNode)
			staticNodes = append(staticNodes, staticNode)
		}
	}

	testSteps := []struct {
		velocity  zounds.Point
		mvNodePos int
	}{
		{zounds.Point{0, 0}, 2},
		{zounds.Point{0, -4}, 2},
		{zounds.Point{0, -2}, 1},
		{zounds.Point{0, -5}, 0},
		{zounds.Point{0, 11}, 2},
		{zounds.Point{0, 4}, 2},
		{zounds.Point{0, 2}, 3},
		{zounds.Point{0, 5}, 4},
	}

	for _, step := range testSteps {
		drawRecorders := make([]drawRecorder, 5)
		staticNodesIndex := 0
		for i := 0; i < len(drawRecorders); i++ {
			if i == step.mvNodePos {
				drawRecorders[i] = movableNode.EXPECT()
			} else {
				drawRecorders[i] = staticNodes[staticNodesIndex].EXPECT()
				staticNodesIndex++
			}
		}

		movableNode.velocity = step.velocity
		s.world.Update()
		checkDrawOrder(s.world, drawRecorders...)
	}
}

type drawRecorder interface {
	Draw() *gomock.Call
}

func checkDrawOrder(w *zounds.World, drawRecorders ...drawRecorder) {
	var prevCall *gomock.Call
	for _, rec := range drawRecorders {
		if prevCall == nil {
			prevCall = rec.Draw()
		} else {
			prevCall = rec.Draw().After(prevCall)
		}
	}
	w.Draw()
}

func TestWorld(t *testing.T) {
	suite.Run(t, new(WorldSuite))
}
