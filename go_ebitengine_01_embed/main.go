package main

import (
	"embed"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/sfnt"
)

//go:embed assets
var assets embed.FS

var (
	helloEmbeddedFile string
	otherEmbeddedFile string
	imageEmbeddedFile *ebiten.Image
	fontEmbeddedFile *sfnt.Font
	mplusNormalFont font.Face
	mplusHudFont font.Face
	soundEmbeddedFile *Sound
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, helloEmbeddedFile + ", " + otherEmbeddedFile)
	text.Draw(screen, "Hello", mplusNormalFont, 20, 80, color.White)
	text.Draw(screen, "World", mplusHudFont, 50, 120, color.White)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(200,200)
	op.GeoM.Scale(1, 1)
	screen.DrawImage(imageEmbeddedFile, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {

	helloEmbeddedFile = string(loadStaticResource(assets, "assets/texts/hello.txt"))
	otherEmbeddedFile = string(loadStaticResource(assets, "assets/texts/other.txt"))
	imageEmbeddedFile, _, _ = ebitenutil.NewImageFromFileSystem(assets,"assets/img/fuel.png")
	loadFonts()
	loadSounds()

	soundEmbeddedFile.Play()


	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle(helloEmbeddedFile)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}

}