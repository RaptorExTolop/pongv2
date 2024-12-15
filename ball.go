package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ball struct {
	X, Y, Radius, Speed int32
	Colour              rl.Color
	dirX, dirY          int32
}

func (b *Ball) update(p *Player) {
	b.X += b.Speed * b.dirX
	b.Y += b.Speed * b.dirY
	if b.checkCollisionWithPlayer(p) {
		fmt.Println("collision with paddle")
		b.dirX *= -1
		b.dirY *= -1
	} else {
		if b.X+b.Radius >= windowWidth {
			b.dirX *= -1
			fmt.Println("reached left wall")
			score += 1
		} else if b.X <= 0 {
			b.dirX *= -1
			fmt.Println("reached right wall")
			score = 0
		}
		if b.Y <= 0 {
			b.dirY *= -1
		} else if b.Y >= windowHeight {
			b.dirY *= -1
		}
	}
}

func (b *Ball) checkCollisionWithPlayer(p *Player) bool {
	topFunc := func() bool {
		return b.X-b.Radius <= p.X && b.Y-b.Radius >= int32(p.Y)
	}
	top := topFunc()
	bottomFunc := func() bool {
		return b.X+b.Radius >= p.X+p.Width && b.Y+b.Radius <= int32(p.Y)+p.Height
	}
	bottom := bottomFunc()

	return top && bottom
}
