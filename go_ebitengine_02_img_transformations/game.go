package main

import (
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)


type Game struct {
	alpha         float64
	degrees       float64
	posX    	  float64
	fuelObjects   []Fuel
	playerObjects []Player
}

func NewGame() *Game {

	g := &Game{
		alpha: 0,
		posX:  0,
	}

	return g
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
	
	g.fuelObjects[0] = ApplyMovement(g.fuelObjects[0], g.posX)
	g.fuelObjects[1] = ApplyMovementFlipX(g.fuelObjects[1], g.posX)
	g.fuelObjects[2] = ApplyMovementAlpha(g.fuelObjects[2], g.posX, g.alpha)
	g.fuelObjects[3] = ApplyMovementRotationAlpha(g.fuelObjects[3], g.posX, g.alpha, g.degrees)
	g.fuelObjects[4] = ApplyMovementWhite(g.fuelObjects[4], g.posX)
	
	g.playerObjects[0] = ApplyMovement(g.playerObjects[0], g.posX + 200)
	g.playerObjects[1] = ApplyMovementFlipX(g.playerObjects[1], g.posX + 200)
	g.playerObjects[2] = ApplyMovementAlpha(g.playerObjects[2], g.posX + 200, g.alpha)
	g.playerObjects[3] = ApplyMovementRotationAlpha(g.playerObjects[3], g.posX + 200, g.alpha, g.degrees)
	g.playerObjects[4] = ApplyMovementWhite(g.playerObjects[4], g.posX + 200)

	return nil
}


func (g *Game) Draw(screen *ebiten.Image) {

	ebitenutil.DebugPrint(screen, strconv.Itoa(int(g.posX)))

	for _, object := range g.fuelObjects {
		object.Draw(screen)
	} 

	for _, object := range g.playerObjects {
		object.Draw(screen)
	} 
	
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return appWidth, appHeight
}


func (g *Game) Init() error {
	
	fuelImage, _, _ := ebitenutil.NewImageFromFileSystem(assets, "assets/img/fuel.png")
	playerImage, _, _ := ebitenutil.NewImageFromFileSystem(assets, "assets/img/player_right.png")

	g.fuelObjects = append(g.fuelObjects,
		NewFuel(65, fuelImage),
		NewFuel(65*2, fuelImage),
		NewFuel(65*3, fuelImage),
		NewFuel(65*4, fuelImage),
		NewFuel(65*5, fuelImage),
	)

	g.playerObjects = append(g.playerObjects,
		NewPlayer(65, playerImage),
		NewPlayer(65*2, playerImage),
		NewPlayer(65*3, playerImage),
		NewPlayer(65*4, playerImage),
		NewPlayer(65*5, playerImage),
	)

	return nil
}

func (g *Game) Exit() error {
	//TODO: finish sounds and music
	return nil
}