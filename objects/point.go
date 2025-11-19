package objects

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/pen787/SentionLevelBuilder/emath"
)

type Point struct {
	ID       uint32
	Layer    uint8
	Position emath.Vector2
	Selected bool
}

func (p *Point) GetLayer() uint8 {
	return p.Layer
}

func (p *Point) SetLayer(layer uint8) {
	p.Layer = layer
}

func (p *Point) SetID(ID uint32) {
	p.ID = ID
}

func (p *Point) GetID() uint32 {
	return p.ID
}

func (p *Point) Ready() error {
	return nil
}

func (p *Point) Exit() error {
	return nil
}

func (p *Point) Render(gs Manager, layer *Layers, screen *ebiten.Image) {
	if p.Selected {
		vector.FillCircle(screen, p.Position.X, p.Position.Y, 8, color.RGBA{255, 255, 0, 255}, false)
	} else {
		vector.FillCircle(screen, p.Position.X, p.Position.Y, 5, color.RGBA{255, 0, 0, 255}, false)
	}
}

func NewPoint(pos emath.Vector2) *Point {
	return &Point{
		Position: pos,
		ID:       0,
		Layer:    0,
		Selected: false,
	}
}
