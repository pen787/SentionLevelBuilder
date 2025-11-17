package objects

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/pen787/SentionLevelBuilder/emath"
)

type Point struct {
	ID       uint32
	Position emath.Vector2
}

func (p *Point) GetPosition() emath.Vector2 {
	return p.Position
}

func (p *Point) SetPosition(vector emath.Vector2) {
	p.Position = vector
}

func (p *Point) SetID(ID uint32) {
	p.ID = ID
}

func (p *Point) GetID() uint32 {
	return p.ID
}

func (p *Point) Render(gs Manager, layer Layer, screen *ebiten.Image) {
	vector.FillCircle(screen, p.Position.X, p.Position.Y, 5, color.RGBA{255, 0, 0, 255}, false)
}

func NewPoint(pos emath.Vector2) *Point {
	return &Point{
		Position: pos,
	}
}
