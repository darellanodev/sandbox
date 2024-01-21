package main

import (
	"github.com/darellanodev/sandbox-go-ebitengine-02-img-transformations/lib"
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

func (o *Object) MoveTo(x float64, y float64) {
	o.x = x
	o.y = y
	o.op = lib.MoveImg(o.x, o.y, o.op)
}

func (o *Object) MoveToX(x float64) {
	o.MoveTo(x, o.y)
}

func (o *Object) Rotate(degrees float64) {
	o.degrees = degrees
	o.op = lib.RotateImg(o.img, o.degrees, o.op)
}

func (o *Object) Alpha(alpha float64) {
	o.alpha = alpha
	o.cm = lib.AlphaImg(o.op, o.cm, o.alpha)
}

func (o *Object) White() {
	o.cm = lib.WhiteImg(o.op, o.cm)
}

func (o *Object) Reset() {
	o.op = &colorm.DrawImageOptions{}
	o.cm = colorm.ColorM{}
}

func (o *Object) FlipX() {
	o.op = lib.FlipXImg(o.img, o.op)
}


func (o *Object) Draw(screen *ebiten.Image) {
	colorm.DrawImage(screen, o.img, o.cm, o.op)
}
