package main

import (
	"embed"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	appWidth = 640
	appHeight = 480
)

//go:embed assets
var assets embed.FS

func main() {
	g := NewGame()

	if err := g.Init(); err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(appWidth, appHeight)
	ebiten.SetWindowTitle("Sandbox")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeDisabled)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

	if err := g.Exit(); err != nil {
		log.Fatal(err)
	}
}