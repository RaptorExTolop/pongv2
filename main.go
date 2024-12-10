package main

import rl "github.com/gen2brain/raylib-go/raylib"

type GameObject interface {
	update()
	draw()
}

type gameObjects []GameObject

var (
	running  bool = true
	gameObjs gameObjects
)

func update() {
	running = !rl.WindowShouldClose()
}

func draw() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)
	rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

	rl.EndDrawing()
}

func init() {
	rl.SetTargetFPS(60)
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
}

func main() {
	defer rl.CloseWindow()
	for running {
		update()
		draw()
	}
}
