package games

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pen787/SentionLevelBuilder/emath"
	"github.com/pen787/SentionLevelBuilder/objects"
)

type Game struct {
	Layers      *objects.Layers
	GameManager objects.Manager
}

func (g *Game) Update() error {
	err := g.GameManager.Update()
	if err != nil {
		return err
	}

	for _, layer := range g.Layers.Data {
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
	for _, layer := range g.Layers.Data {
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
	IDGenerator := emath.NewIDGenerator()
	layers := objects.MakeLayer(IDGenerator)
	return &Game{
		Layers:      layers,
		GameManager: objects.NewGameManager(layers),
	}
}
