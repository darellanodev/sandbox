package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func getImageCenter(img *ebiten.Image) (float64, float64) {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	halfW := float64(width / 2)
	halfH := float64(height / 2)

	return halfW, halfH 
}

func drawRotateImage(screen *ebiten.Image, img *ebiten.Image, posX float64, posY float64, degrees float64) {
	
	halfW, halfH := getImageCenter(img)
	
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(degrees * math.Pi / 180.0)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(posX, posY)
	screen.DrawImage(img, op)
}

func drawNormalImage(screen *ebiten.Image, img *ebiten.Image, posX float64, posY float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 1)
	op.GeoM.Translate(posX, posY)
	screen.DrawImage(img, op)
}

func drawHorizontalFlippedImage(screen *ebiten.Image, img *ebiten.Image, imageWidth float64, posX float64, posY float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(-1, 1)
	op.GeoM.Translate(posX+imageWidth, posY)
	screen.DrawImage(img, op)
}