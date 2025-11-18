package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/pen787/SentionLevelBuilder/emath"
)

type Manager interface {
	Update() error
}

// Implement Object
type GameManager struct {
	Layers *Layers
}

func (gm *GameManager) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		ok := gm.Layers.AddObject(1, NewPoint(emath.GetCursor()))
	}

	return nil
}

func NewGameManager(layers *Layers) *GameManager {
	return &GameManager{
		Layers: layers,
	}
}
