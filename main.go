package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	X      int32
	Y      int32
	Width  int32
	Height int32
	colour rl.Color
}

var (
	// game setup
	running                        bool
	dt                             float32
	windowHeight, windowWidth, fps int32

	player = Player{}
)

func input() {}

func update() {
	fps = rl.GetFPS()
	running = !rl.WindowShouldClose()
	dt = rl.GetFrameTime()
}

func draw() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.Green)

	rl.DrawRectangle(player.X, player.Y, player.Width, player.Height, player.colour)

	/*
		Debugging statments
	*/
	rl.DrawText(fmt.Sprint("FPS: ", fps), 0, 0, 24, rl.White)
	rl.DrawText(fmt.Sprint("DT:  ", dt), 0, 28, 24, rl.White)

	rl.EndDrawing()
}

func init() {
	windowHeight = 720
	windowWidth = 1280

	rl.InitWindow(windowWidth, windowHeight, "Pong")
	running = true
	player = Player{5, (windowHeight/2 - player.Height/2), 60, 180, rl.RayWhite}
}

func quit() {
	rl.CloseWindow()
}

func main() {
	defer quit()

	for running {
		input()
		update()
		draw()
	}
}
