package lib

import (
	"embed"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func getImageCenter(img *ebiten.Image) (float64, float64) {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	halfW := float64(width / 2)
	halfH := float64(height / 2)

	return halfW, halfH 
}

func RotateImg(img *ebiten.Image, degrees float64, op *colorm.DrawImageOptions) *colorm.DrawImageOptions {
	halfW, halfH := getImageCenter(img)
	
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(degrees * math.Pi / 180.0)
	op.GeoM.Translate(halfW, halfH)

	return op
}

func MoveImg(posX float64, posY float64, op *colorm.DrawImageOptions) *colorm.DrawImageOptions {
	op.GeoM.Translate(posX, posY)
	return op
}

func AlphaImg(op *colorm.DrawImageOptions, cm colorm.ColorM, alpha float64) colorm.ColorM {

	cm.Scale(1.0, 1.0, 1.0, alpha)

	return cm
}

func FlipXImg(img *ebiten.Image, op *colorm.DrawImageOptions) *colorm.DrawImageOptions {

	width := float64(img.Bounds().Dx())

	op.GeoM.Translate(-width, 0)
	op.GeoM.Scale(-1, 1)
	
	return op
}

func WhiteImg(op *colorm.DrawImageOptions, cm colorm.ColorM) colorm.ColorM {

	cm.Translate(1.0, 1.0, 1.0, 0.0)

	return cm
}

func LoadImage(filesystem embed.FS , path string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFileSystem(filesystem, path)

	if err != nil {
		log.Fatal(err)
	}

	return img
}
