package main

import (
	"embed"
	_ "image/png"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	appWidth = 640
	appHeight = 480

	offsetY = 100
)

//go:embed assets
var assets embed.FS

var (
	imageEmbeddedFile *ebiten.Image
	imagePosX  int
	imagePosY  int
	imageWidth int
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}



func (g *Game) Draw(screen *ebiten.Image) {

	ebitenutil.DebugPrint(screen, strconv.Itoa(imagePosX) + ", " + strconv.Itoa(imagePosY))

	if (imagePosX < appWidth) {
		imagePosX++
	} else {
		imagePosX = 0
	}
	drawNormalImage(screen, imageEmbeddedFile, imagePosX, imagePosY)
	drawHorizontalFlippedImage(screen, imageEmbeddedFile, imageWidth, imagePosX, imagePosY + offsetY)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return appWidth, appHeight
}

func main() {

	// initialization
	imageEmbeddedFile, _, _ = ebitenutil.NewImageFromFileSystem(assets,"assets/img/fuel.png")

	imagePosX = 0
	imagePosY = 0
	imageWidth = 32

	ebiten.SetWindowSize(appWidth, appHeight) // Note that the values are equals to layout function
	ebiten.SetWindowTitle("flip image")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}