package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ball struct {
	X, Y, Radius int32
	Colour       rl.Color
}

func (b *Ball) update() int32 {
	return 0 //int32(math.Round(float64(dt)) * 1)
}
