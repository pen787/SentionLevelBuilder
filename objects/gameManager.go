package objects

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/pen787/SentionLevelBuilder/emath"
)

type STATE_TOOL uint8

const (
	STATE_TOOL_SELECT STATE_TOOL = iota
	STATE_TOOL_ADD
	STATE_TOOL_CONNECT
	STATE_TOOL_MOVE
)

type Manager interface {
	Ready() error
	Update() error
}

type GameManager struct {
	Layers         *Layers
	SelectedPoints map[*Point]struct{}
	Selected_tool  STATE_TOOL
}

func (gm *GameManager) SelectPoint(p *Point) {
	p.Select()
	gm.SelectedPoints[p] = struct{}{}

}

func (gm *GameManager) DeselectPoint(p *Point) {
	if _, ok := gm.SelectedPoints[p]; ok {
		p.Deselect()
		delete(gm.SelectedPoints, p)
	}
}

func (gm *GameManager) DeselectPoints() {
	for i, _ := range gm.SelectedPoints {
		gm.DeselectPoint(i)
	}
}

func (gm *GameManager) ClickPoint() {
	for _, p := range gm.Layers.Data[1] {
		if p, ok := p.(*Point); ok {
			if p.Position.DistanceTo(emath.GetCursor()) < 10 {
				if _, ok := gm.SelectedPoints[p]; ok {
					p.Deselect()
					gm.DeselectPoint(p)
				} else {
					gm.SelectPoint(p)
					p.Select()
				}
				break
			}
		}
	}
}

func (gm *GameManager) AddConnectPoint() {

}

func (gm *GameManager) AddPointAtCursor() {
	gm.DeselectPoints()

	p := NewPoint(emath.GetCursor())
	gm.Layers.AddObject(1, p)
	gm.SelectPoint(p)
}

func (gm *GameManager) Ready() error {
	return nil
}

func (gm *GameManager) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {

		switch gm.Selected_tool {
		case STATE_TOOL_SELECT:
			gm.ClickPoint()
		case STATE_TOOL_ADD:
			gm.AddPointAtCursor()
		case STATE_TOOL_CONNECT:
			gm.DeselectPoints()
			gm.AddConnectPoint()
		}

	}

	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		gm.Selected_tool = ((gm.Selected_tool - 1) + 3) % 3
		log.Printf("Tool : %d\n", gm.Selected_tool)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		gm.Selected_tool = (gm.Selected_tool + 1) % 3
		log.Printf("Tool : %d\n", gm.Selected_tool)
	}

	return nil
}

func NewGameManager(layers *Layers) *GameManager {
	return &GameManager{
		Layers:         layers,
		SelectedPoints: make(map[*Point]struct{}),
		Selected_tool:  STATE_TOOL_SELECT,
	}
}
