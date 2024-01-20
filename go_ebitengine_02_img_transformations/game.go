package main

import (
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	image   *ebiten.Image
	alpha   float64
	degrees   float64
	posX    float64
	objects []Object
}

func NewGame() *Game {

	g := &Game{
		alpha: 0,
		posX:  0,
	}

	return g
}

func (g *Game) applyMovement(obj Object, posX float64) Object {
	obj.Reset()
	obj.MoveTo(posX, obj.y)

	return obj
}

func (g *Game) applyMovementFlipX(obj Object, posX float64) Object {
	obj.Reset()
	obj.FlipX() // the order is important. If you want to flip the object do it before moving it.
	obj.MoveTo(posX, obj.y)

	return obj
}

func (g *Game) applyMovementAlpha(obj Object, posX float64) Object {
	obj.Reset()
	obj.MoveTo(posX, obj.y)
	obj.Alpha(g.alpha)

	return obj
}

func (g *Game) applyMovementRotationAlpha(obj Object, posX float64) Object {
	obj.Reset()
	obj.Rotate(g.degrees) // the order is important. If you want to rotate the object do it before moving it.
	obj.MoveTo(posX, obj.y)
	obj.Alpha(g.alpha)

	return obj
}

func (g *Game) applyMovementWhite(obj Object, posX float64) Object {
	obj.Reset()
	obj.White()
	obj.MoveTo(posX, obj.y)

	return obj
}

func (g *Game) updatePosX() {
	if g.posX < appWidth {
		g.posX++
	} else {
		g.posX = 0
	}
}

func (g *Game) updateDegrees() {
	if g.degrees < 360 {
		g.degrees++
	} else {
		g.degrees = 0
	}
}

func (g* Game) updateAlpha() {
	if g.alpha > 0 {
		g.alpha -= 0.01
	} else {
		g.alpha = 1
	}
}

func (g *Game) Update() error {

	g.updatePosX()
	g.updateDegrees()
	g.updateAlpha()
	
	g.objects[0] = g.applyMovement(g.objects[0], g.posX)
	g.objects[1] = g.applyMovementFlipX(g.objects[1], g.posX)
	g.objects[2] = g.applyMovementAlpha(g.objects[2], g.posX)
	g.objects[3] = g.applyMovementRotationAlpha(g.objects[3], g.posX)
	g.objects[4] = g.applyMovementWhite(g.objects[4], g.posX)

	return nil
}


func (g *Game) Draw(screen *ebiten.Image) {

	ebitenutil.DebugPrint(screen, strconv.Itoa(int(g.posX)))

	for _, object := range g.objects {
		object.Draw(screen)
	} 
	
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return appWidth, appHeight
}


func (g *Game) Init() error {
	
	g.image, _, _ = ebitenutil.NewImageFromFileSystem(assets, "assets/img/fuel.png")

	g.objects = append(g.objects,
		NewObject(50, g.image),
		NewObject(100, g.image),
		NewObject(150, g.image),
		NewObject(200, g.image),
		NewObject(250, g.image),
	)

	return nil
}

func (g *Game) Exit() error {
	//TODO: finish sounds and music
	return nil
}