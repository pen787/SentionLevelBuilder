package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Updateable interface {
	Update(gs Manager, layer *Layers) error
}

type Renderable interface {
	Render(gs Manager, layer *Layers, screen *ebiten.Image)
}

type Object interface {
	SetID(id uint32)
	GetID() uint32

	SetLayer(layer uint8)
	GetLayer() uint8

	Ready() error
	Exit() error
}
