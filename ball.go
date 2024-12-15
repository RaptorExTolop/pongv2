package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ball struct {
	X, Y, Radius, Speed int32
	Colour              rl.Color
	dir                 int32
}

func (b *Ball) update(p *Player) {
	if b.X-b.Radius <= p.X && b.X+b.Radius >= p.X {
		fmt.Println("collision with paddle")
		b.dir *= 1
	}

	if b.X+b.Radius >= windowWidth {
		b.dir *= -1
		fmt.Println("reached left wall")
		score += 1
	} else if b.X <= 0 {
		b.dir *= -1
		fmt.Println("reached right wall")
		score = 0
	}
	b.X += b.Speed * b.dir
}
