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

	offsetY = 70
)

//go:embed assets
var assets embed.FS

var (
	imageEmbeddedFile *ebiten.Image
	imagePosX  float64
	imagePosY  float64
	imageWidth float64
	degrees	   float64
	newOffsetY float64
)

type Game struct{}

func (g *Game) Update() error {

	if imagePosX < appWidth {
		imagePosX++
	} else {
		imagePosX = 0
	}

	if degrees < 360 {
		degrees++
	} else {
		degrees = 0
	}

	return nil
}

func getY(row int, initialY float64, offset float64) float64 {
	return initialY + (offset * float64(row))
}

func (g *Game) Draw(screen *ebiten.Image) {

	ebitenutil.DebugPrint(screen, strconv.Itoa(int(imagePosX)) + ", " + strconv.Itoa(int(imagePosY)))

	newOffsetY = getY(1, imagePosY, offsetY)
	drawNormalImage(screen, imageEmbeddedFile, imagePosX, newOffsetY)

	newOffsetY = getY(2, imagePosY, offsetY)
	drawHorizontalFlippedImage(screen, imageEmbeddedFile, imageWidth, imagePosX, newOffsetY)

	newOffsetY = getY(3, imagePosY, offsetY)
	drawRotateImage(screen, imageEmbeddedFile, imagePosX, newOffsetY, degrees)

	newOffsetY = getY(4, imagePosY, offsetY)
	drawWhiteImage(screen, imageEmbeddedFile, imagePosX, newOffsetY)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return appWidth, appHeight
}

func main() {

	// initialization
	imageEmbeddedFile, _, _ = ebitenutil.NewImageFromFileSystem(assets,"assets/img/fuel.png")

	imagePosX  = 0
	imagePosY  = 0
	imageWidth = 32
	degrees    = 0

	ebiten.SetWindowSize(appWidth, appHeight) // Note that the values are equals to layout function
	ebiten.SetWindowTitle("flip image")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}