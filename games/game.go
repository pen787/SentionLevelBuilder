package games

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pen787/SentionLevelBuilder/emath"
	"github.com/pen787/SentionLevelBuilder/objects"
)

type Game struct {
	Layers      objects.Layer
	GameManager objects.Manager
	IDGenerator *emath.IDGenerator
}

func (g *Game) AddObject(floor uint16, obj objects.Object) {
	obj.SetID(g.IDGenerator.Next())
	if g.Layers[floor] == nil {
		g.Layers[floor] = make(objects.IDObjects)
	}
	g.Layers[floor][obj.GetID()] = obj
}

func (g *Game) RemoveObject(floor uint16, obj objects.Object) bool {
	_, ok := g.Layers[floor][obj.GetID()]
	if ok {
		g.Layers[floor][obj.GetID()] = nil
		return true
	}

	return false
}

func (g *Game) Update() error {
	err := g.GameManager.Update(g.Layers, g.AddObject, g.RemoveObject)
	if err != nil {
		return err
	}

	for _, layer := range g.Layers {
		for _, obj := range layer {
			got, ok := obj.(objects.Updateable)
			if ok {
				err := got.Update(g.GameManager, g.Layers)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, layer := range g.Layers {
		for _, obj := range layer {
			got, ok := obj.(objects.Renderable)
			if ok {
				got.Render(g.GameManager, g.Layers, screen)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func NewGame() *Game {
	return &Game{
		Layers:      make(objects.Layer, 10, 10),
		GameManager: objects.NewGameManager(),
		IDGenerator: emath.NewIDGenerator(),
	}
}
