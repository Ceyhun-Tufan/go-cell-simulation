package main

import rl "github.com/gen2brain/raylib-go/raylib"

// classes

// implementations

func init() {
	rl.SetConfigFlags(rl.FlagMsaa4xHint)
}

func main() {
	rl.InitWindow(800, 600, "Cell Simulation")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Hello, Raylib!", 190, 200, 20, rl.DarkGray)

		rl.EndDrawing()
	}
}
