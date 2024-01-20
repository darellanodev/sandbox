package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
)

type Player struct {
	*Object
}

func NewPlayer(posY float64, image *ebiten.Image) Player {

	return Player {
		&Object{
			x:       0,
			y:       posY,
			op:      &colorm.DrawImageOptions{},
			cm:      colorm.ColorM{},
			img:     image,
			degrees: 0,
			alpha:   1,
		},
	}	
}