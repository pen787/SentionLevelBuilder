package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pen787/SentionLevelBuilder/games"
)

const (
	window_width  = 840
	widnow_height = 640
)

func main() {
	ebiten.SetWindowSize(window_width, widnow_height)
	ebiten.SetWindowTitle("Sention Level Builder")

	game := games.NewGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
