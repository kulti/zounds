// +build examples

package main

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"

	"github.com/kulti/zounds"
	zebiten "github.com/kulti/zounds/ebiten"
)

func main() {
	ebiten.SetWindowSize(640, 480)

	world := zounds.NewWorld()
	game := zebiten.NewGame(world, image.Rect(0, 0, 640, 480))

	img, _ := ebiten.NewImage(640, 480, ebiten.FilterDefault)
	img.Fill(color.White)
	world.SetBackground(game.NewImageNode(img, zounds.Point{}))

	img, _ = ebiten.NewImage(100, 100, ebiten.FilterDefault)
	img.Fill(color.RGBA{255, 0, 0, 255})
	node := game.NewImageNode(img, zounds.Point{100, 50})
	world.AddStaticNode(node)

	img, _ = ebiten.NewImage(20, 20, ebiten.FilterDefault)
	img.Fill(color.RGBA{0, 255, 0, 255})
	player := game.NewPlayer(img, zounds.Point{50, 150}, 240)
	world.AddMovableNode(player)

	ebiten.RunGame(game)
}
