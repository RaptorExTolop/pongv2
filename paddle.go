package main

import rl "github.com/gen2brain/raylib-go/raylib"

type paddle struct {
	PosX, PosY    int32
	Width, Hieght int32
	Colour        rl.Color
	Hitbox        HitBox
}

func (p *paddle) init(posX, posY, width, height int32, colour rl.Color) {
	p.PosX = posX
	p.PosY = posY
	p.Width = width
	p.Hieght = height
	p.Colour = colour

	p.Hitbox = HitBox{p.PosX, p.PosX + width, p.PosY, p.PosY + p.Hieght}
}
