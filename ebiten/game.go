package zebiten

import (
	"image"

	"github.com/hajimehoshi/ebiten"

	"github.com/kulti/zounds"
)

// Game implements ebiten.Game interface and glues zounds and ebiten.
type Game struct {
	world  *zounds.World
	screen *screen
	layout image.Rectangle
}

// NewGame creates a new instance of Game.
func NewGame(world *zounds.World, layout image.Rectangle) *Game {
	return &Game{
		world:  world,
		screen: newScreen(),
		layout: layout,
	}
}

// Update implements ebiten.Game Update method.
func (g *Game) Update(_ *ebiten.Image) error {
	g.world.Update()
	return nil
}

// Draw implements ebiten.Game Draw method.
func (g *Game) Draw(screen *ebiten.Image) {
	g.screen.UpdateImage(screen)
	g.world.Draw()
}

// Layout implements ebiten.Game Layout method.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.layout.Dx(), g.layout.Dy()
}
