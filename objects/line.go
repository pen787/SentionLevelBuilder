package objects

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Line struct {
	ID    uint32
	Layer uint8
	P1    *Point
	P2    *Point
}

func (l *Line) GetLayer() uint8 {
	return l.Layer
}

func (l *Line) SetLayer(layer uint8) {
	l.Layer = layer
}

func (l *Line) SetID(ID uint32) {
	l.ID = ID
}

func (l *Line) GetID() uint32 {
	return l.ID
}

func (l *Line) Ready() error {
	return nil
}

func (l *Line) Exit() error {
	return nil
}

func (l *Line) Render(gs Manager, layer *Layers, screen *ebiten.Image) {
	vector.StrokeLine(screen, l.P1.Position.X, l.P1.Position.Y, l.P2.Position.X, l.P2.Position.Y, 2, color.RGBA{255, 255, 255, 255}, true)
}

func NewLine(P1 *Point, P2 *Point) *Line {
	return &Line{
		P1:    P1,
		P2:    P2,
		ID:    0,
		Layer: 0,
	}
}
