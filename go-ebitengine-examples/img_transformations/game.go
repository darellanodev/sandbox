package main

import (
	"image/color"
	"strconv"

	"github.com/darellanodev/sandbox-go-ebitengine-02-img-transformations/lib"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)


type Game struct {
	alpha          float64
	degrees        float64
	fuelObjectsX   float64
	playerObjectsX float64
	fuelObjects    []Fuel
	playerObjects  []Player
	font		   font.Face
}

func NewGame() *Game {

	g := &Game{
		alpha: 0,
		fuelObjectsX:  0,
		playerObjectsX:  200,
	}

	return g
}

func (g *Game) updatePosX() {
	if g.fuelObjectsX < appWidth {
		g.fuelObjectsX++
	} else {
		g.fuelObjectsX = 0
	}
	if g.playerObjectsX < appWidth {
		g.playerObjectsX++
	} else {
		g.playerObjectsX = 0
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

	
	g.fuelObjects[0] = ApplyMovement(g.fuelObjects[0], g.fuelObjectsX)
	g.fuelObjects[1] = ApplyMovementFlipX(g.fuelObjects[1], g.fuelObjectsX)
	g.fuelObjects[2] = ApplyMovementAlpha(g.fuelObjects[2], g.fuelObjectsX, g.alpha)
	g.fuelObjects[3] = ApplyMovementRotationAlpha(g.fuelObjects[3], g.fuelObjectsX, g.alpha, g.degrees)
	g.fuelObjects[4] = ApplyMovementWhite(g.fuelObjects[4], g.fuelObjectsX)
	
	g.playerObjects[0] = ApplyMovement(g.playerObjects[0], g.playerObjectsX)
	g.playerObjects[1] = ApplyMovementFlipX(g.playerObjects[1], g.playerObjectsX)
	g.playerObjects[2] = ApplyMovementAlpha(g.playerObjects[2], g.playerObjectsX, g.alpha)
	g.playerObjects[3] = ApplyMovementRotationAlpha(g.playerObjects[3], g.playerObjectsX, g.alpha, g.degrees)
	g.playerObjects[4] = ApplyMovementWhite(g.playerObjects[4], g.playerObjectsX)

	return nil
}


func (g *Game) Draw(screen *ebiten.Image) {

	ebitenutil.DebugPrint(screen, strconv.Itoa(int(g.fuelObjectsX)))
	text.Draw(screen, "image transformations", g.font, 50, 30, color.White)


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
	
	fuelImage := lib.LoadImage(assets, "assets/img/fuel.png")
	playerImage := lib.LoadImage(assets, "assets/img/player_right.png")
	g.font = lib.LoadFont(assets, "assets/fonts/pressstart2p.ttf")

	totalObjects := 5
	
	for i := 0; i < totalObjects; i++ {
		g.fuelObjects = append(g.fuelObjects, NewFuel(float64(65*(i+1)), fuelImage))
	}

	for i := 0; i < totalObjects; i++ {
		g.playerObjects = append(g.playerObjects, NewPlayer(float64(65*(i+1)), playerImage))
	}
	
	return nil
}

func (g *Game) Exit() error {
	//TODO: finish sounds and music
	return nil
}