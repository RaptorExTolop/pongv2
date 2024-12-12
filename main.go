package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	running bool

	dt float32

	windowHeight, windowWidth int32
)

func input() {}

func update() {
	running = !rl.WindowShouldClose()
	dt = rl.GetFrameTime()
}

func draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Green)
	rl.DrawText(fmt.Sprint("FPS: ", rl.GetFPS()), 0, 0, 24, rl.White)
	rl.DrawText(fmt.Sprint("DT:  ", dt), 0, 28, 24, rl.White)

	rl.EndDrawing()
}

func init() {
	windowHeight = 720
	windowWidth = 1280

	rl.InitWindow(windowWidth, windowHeight, "Pong")
	running = true
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
