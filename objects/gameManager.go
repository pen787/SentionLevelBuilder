package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/pen787/SentionLevelBuilder/emath"
)

type Manager interface {
	Ready() error
	Update() error
}

type GameManager struct {
	Layers   *Layers
	Selected *Point
}

func (gm *GameManager) Ready() error {
	return nil
}

func (gm *GameManager) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		if gm.Selected != nil {
			gm.Selected.Selected = false
			gm.Selected = nil
		}

		for _, p := range gm.Layers.Data[1] {
			if p, ok := p.(*Point); ok {
				if p.Position.DistanceTo(emath.GetCursor()) < 10 {
					gm.Selected = p
					p.Selected = true
					break
				} else {
					gm.Selected = nil
				}
			}
		}

		if gm.Selected == nil {
			gm.Layers.AddObject(1, NewPoint(emath.GetCursor()))
		}
	}

	return nil
}

func NewGameManager(layers *Layers) *GameManager {
	return &GameManager{
		Layers: layers,
	}
}
