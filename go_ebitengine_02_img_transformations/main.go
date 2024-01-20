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
)

//go:embed assets
var assets embed.FS

var (
	image *ebiten.Image
	posX  float64
	imagePosY  float64

	degrees	   float64
	alpha	   float64
	objects	   []Object
	
)

type Game struct{}

func applyMovement(obj Object, posX float64) Object {
	obj.Reset()
	obj.MoveTo(posX, obj.y)

	return obj
}

func applyMovementFlipX(obj Object, posX float64) Object {
	obj.Reset()
	obj.FlipX() // the order is important. If you want to flip the object do it before moving it.
	obj.MoveTo(posX, obj.y)

	return obj
}

func applyMovementAlpha(obj Object, posX float64) Object {
	obj.Reset()
	obj.MoveTo(posX, obj.y)
	obj.Alpha(alpha)

	return obj
}

func applyMovementRotationAlpha(obj Object, posX float64) Object {
	obj.Reset()
	obj.Rotate(degrees) // the order is important. If you want to rotate the object do it before moving it.
	obj.MoveTo(posX, obj.y)
	obj.Alpha(alpha)

	return obj
}

func applyMovementWhite(obj Object, posX float64) Object {
	obj.Reset()
	obj.White()
	obj.MoveTo(posX, obj.y)

	return obj
}

func (g *Game) updatePosX() {
	if posX < appWidth {
		posX++
	} else {
		posX = 0
	}
}

func (g *Game) updateDegrees() {
	if degrees < 360 {
		degrees++
	} else {
		degrees = 0
	}
}

func (g* Game) updateAlpha() {
	if alpha > 0 {
		alpha -= 0.01
	} else {
		alpha = 1
	}
}

func (g *Game) Update() error {

	g.updatePosX()
	g.updateDegrees()
	g.updateAlpha()
	
	objects[0] = applyMovement(objects[0], posX)
	objects[1] = applyMovementFlipX(objects[1], posX)
	objects[2] = applyMovementAlpha(objects[2], posX)
	objects[3] = applyMovementRotationAlpha(objects[3], posX)
	objects[4] = applyMovementWhite(objects[4], posX)

	return nil
}


func (g *Game) Draw(screen *ebiten.Image) {

	ebitenutil.DebugPrint(screen, strconv.Itoa(int(posX)) + ", " + strconv.Itoa(int(imagePosY)))

	for _, object := range objects {
		object.Draw(screen)
	} 
	
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return appWidth, appHeight
}

func main() {

	image, _, _ = ebitenutil.NewImageFromFileSystem(assets,"assets/img/fuel.png")

	degrees    = 0
	alpha	   = 1

	objects = append(objects, 
		NewObject(50), 
		NewObject(100),
		NewObject(150),
		NewObject(200),
		NewObject(250),
	)

	ebiten.SetWindowSize(appWidth, appHeight) // Note that the values are equals to layout function
	ebiten.SetWindowTitle("flip image")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}