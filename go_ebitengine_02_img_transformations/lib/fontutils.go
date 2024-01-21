package lib

import (
	"embed"
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)


func LoadFont(filesystem embed.FS, path string) font.Face {

	const dpi = 72
	
	fontBytes, err := LoadStaticResource(filesystem, path)

	if err != nil {
		log.Fatal(err)
	}

	tttf, err := opentype.Parse(fontBytes)

	if err != nil {
		log.Fatal(err)
	}

	font, err := opentype.NewFace(tttf, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}
	
	return font
}

