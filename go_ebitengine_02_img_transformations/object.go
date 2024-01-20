package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
)

type Object struct {
	img 	*ebiten.Image
	x   	float64
	y   	float64
	degrees float64
	op  	*colorm.DrawImageOptions
	cm  	colorm.ColorM
	alpha	float64
}

func NewObject(posY float64, image *ebiten.Image) Object {

	return Object{
		x: 0,
		y: posY,
		op: &colorm.DrawImageOptions{},
		cm: colorm.ColorM{},
		img: image,
		degrees: 0,
		alpha: 1,
	}
}

func (o *Object) MoveTo(x float64, y float64) {
	o.x = x
	o.y = y
	o.op = moveImg(o.x, o.y, o.op)
}

func (o *Object) Rotate(degrees float64) {
	o.degrees = degrees
	o.op = rotateImg(o.img, o.degrees, o.op)
}

func (o *Object) Alpha(alpha float64) {
	o.alpha = alpha
	o.cm = alphaImg(o.op, o.cm, o.alpha)
}

func (o *Object) White() {
	o.cm = whiteImg(o.op, o.cm)
}


// func (o *Object) Update() {
	
// }

func (o *Object) Reset() {
	o.op = &colorm.DrawImageOptions{}
	o.cm = colorm.ColorM{}
}

func (o *Object) FlipX() {
	o.op = flipXImg(o.img, o.op)
}


func (o *Object) Draw(screen *ebiten.Image) {
	colorm.DrawImage(screen, o.img, o.cm, o.op)
}
