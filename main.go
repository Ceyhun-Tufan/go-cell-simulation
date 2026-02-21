package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// consts

const SCREEN_HEIGHT int32 = 600
const SCREEN_WIDHT int32 = 800
const CELL_COUNT int = 100

var cells []Cell

// classes

// implementations

func init() {
	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	initCells()
}

func main() {
	rl.InitWindow(SCREEN_WIDHT, SCREEN_HEIGHT, "Cell Simulation")
	defer rl.CloseWindow()

	fmt.Print(cells)

	var pause bool

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		if rl.IsKeyPressed(rl.KeySpace) {
			pause = !pause
		}

		if !pause {
			// pause logic
		}

		// draw
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		for i := range cells {
			tmp := &cells[i]
			rl.DrawCircleV(tmp.pos, tmp.radius, tmp.color)
		}

		rl.EndDrawing()
	}
}

func initCells() {
	for range CELL_COUNT {
		cells = append(cells, NewBlueCell())
	}
}
