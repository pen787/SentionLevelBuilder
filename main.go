package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pen787/SentionLevelBuilder/emath"
	"github.com/pen787/SentionLevelBuilder/games"
	"github.com/pen787/SentionLevelBuilder/objects"
)

const (
	window_width  = 840
	widnow_height = 640
)

func main() {
	ebiten.SetWindowSize(window_width, widnow_height)
	ebiten.SetWindowTitle("Sention Level Builder")

	game := games.NewGame()
	ok := game.Layers.AddObject(0, objects.NewPoint(emath.NewVector2(20, 20)))
	log.Println(ok)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
