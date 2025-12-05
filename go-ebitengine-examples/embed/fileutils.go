package main

import (
	"embed"
	"io"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// for loading fonts (.ttf) and text files (.txt)
func loadStaticResource(filesystem embed.FS, path string) []byte {
	file, err := assets.Open(path)
	if err != nil {
		panic(err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return data
}

func loadFonts() {

	fontBytes := loadStaticResource(assets, "assets/fonts/pressstart2p.ttf")

	fontEmbeddedFile, _ = opentype.Parse(fontBytes)

	const dpi = 72

	mplusNormalFont, _ = opentype.NewFace(fontEmbeddedFile, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})

	mplusHudFont, _ = opentype.NewFace(fontEmbeddedFile, &opentype.FaceOptions{
		Size:    8,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})

}