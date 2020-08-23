package zebiten

import "github.com/hajimehoshi/ebiten"

type screen struct {
	img *ebiten.Image
}

func newScreen() *screen {
	return &screen{}
}

func (s *screen) DrawImage(img *ebiten.Image, options *ebiten.DrawImageOptions) {
	_ = s.img.DrawImage(img, options)
}

func (s *screen) UpdateImage(img *ebiten.Image) {
	s.img = img
}
