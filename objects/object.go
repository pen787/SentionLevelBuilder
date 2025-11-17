package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type IDObjects map[uint32]Object
type Layer []IDObjects

type Updateable interface {
	Update(gs Manager, layer Layer) error
}

type Renderable interface {
	Render(gs Manager, layer Layer, screen *ebiten.Image)
}

type Object interface {
	SetID(id uint32)
	GetID() uint32
}

type Node interface {
	Updateable
	Renderable
}
