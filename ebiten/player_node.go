package zebiten

import (
	"time"

	"github.com/hajimehoshi/ebiten"

	"github.com/kulti/zounds"
)

// Player contains movable image and handles user input to manage velocity.
type Player struct {
	zounds.FixedSizeNode
	screen       *screen
	img          *ebiten.Image
	velocity     zounds.Point
	baseVelocity float64
}

// NewPlayer creates a new Player instance.
func (g *Game) NewPlayer(img *ebiten.Image, pos zounds.Point, baseVelocity float64) *Player {
	return &Player{
		screen:        g.screen,
		img:           img,
		baseVelocity:  baseVelocity,
		FixedSizeNode: zounds.NewFixedSizeNode(zounds.RectFromImageRect(img.Bounds()).Add(pos)),
	}
}

// Draw dreaw player images.
func (n *Player) Draw() {
	pos := n.Bounds().Min
	drawOpts := &ebiten.DrawImageOptions{}
	drawOpts.GeoM.Translate(pos.X, pos.Y)

	n.screen.DrawImage(n.img, drawOpts)
}

// Update handles user input and updates velocity.
func (n *Player) Update(delta time.Duration) {
	v := n.baseVelocity * float64(delta) / float64(time.Second)

	switch {
	case ebiten.IsKeyPressed(ebiten.KeyUp):
		n.velocity = zounds.Point{X: 0, Y: -v}
	case ebiten.IsKeyPressed(ebiten.KeyDown):
		n.velocity = zounds.Point{X: 0, Y: v}
	case ebiten.IsKeyPressed(ebiten.KeyLeft):
		n.velocity = zounds.Point{X: -v, Y: 0}
	case ebiten.IsKeyPressed(ebiten.KeyRight):
		n.velocity = zounds.Point{X: v, Y: 0}
	default:
		n.velocity = zounds.Point{}
	}
}

// Velocity returns current player's velocity.
func (n *Player) Velocity() zounds.Point {
	return n.velocity
}
