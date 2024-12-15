package main

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	X      int32
	Y      float32
	Width  int32
	Height int32
	colour rl.Color
	dir    int32
	speed  int32
}

var (
	// game setup
	running                        bool
	dt                             float32
	windowHeight, windowWidth, fps int32

	player       = Player{}
	ball         = Ball{}
	score  int32 = 0
)

func input() {
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		player.dir -= 1
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		player.dir += 1
	}
	player.dir = int32(rl.Clamp(float32(player.dir), -1, 1))
	//fmt.Println("Dir.Y:", player.dir)
}

func update() {
	fps = rl.GetFPS()
	running = !rl.WindowShouldClose()
	dt = rl.GetFrameTime()
	player.Y += float32(player.speed*player.dir) * dt
	player.dir = 0
	ball.update(&player)
	//fmt.Println(math.Round(float64(dt)))
	//fmt.Println((float32(player.speed*player.dir) * dt))
}

func draw() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.Green)

	rl.DrawRectangle(player.X, int32(math.Round(float64(player.Y))), player.Width, player.Height, player.colour)

	rl.DrawCircle(ball.X, ball.Y, float32(ball.Radius), ball.Colour)

	/*
		Debugging statments
	*/
	rl.DrawText(fmt.Sprint("FPS: ", fps), 0, 0, 24, rl.White)
	rl.DrawText(fmt.Sprint("DT:  ", dt), 0, 28, 24, rl.White)
	rl.DrawText(fmt.Sprint("SCORE: ", score), 0, 56, 24, rl.White)

	rl.EndDrawing()
}

func init() {
	windowHeight = 720
	windowWidth = 1280

	rl.InitWindow(windowWidth, windowHeight, "Pong")
	rl.SetTargetFPS(60)
	running = true
	player = Player{5, 0, 30, 150, rl.RayWhite, 0, 350}
	player.Y = float32((windowHeight / 2) - (player.Height / 2))

	ball = Ball{0, 0, 25, 10, rl.RayWhite, 1, 1}
	ball.X = windowWidth / 2
	ball.Y = windowHeight/2 - ball.Radius
}

func quit() {
	rl.CloseWindow()
	fmt.Println("ended")
}

func main() {
	defer quit()

	fmt.Println("started")

	for running {
		input()
		update()
		draw()
	}
}
