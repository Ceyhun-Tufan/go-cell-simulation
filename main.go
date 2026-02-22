package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// consts

const SCREEN_HEIGHT int32 = 600
const SCREEN_WIDTH int32 = 800
const CELL_COUNT int = 100
const GRIDCELL_SIZE int = 20

var cells []Cell
var grid *Grid

// classes

// implementations

func init() {
	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	grid = NewGrid(int(SCREEN_WIDTH), int(SCREEN_HEIGHT), GRIDCELL_SIZE)
	initCells()
	grid.insertCellsToGrid(cells)
}

func main() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "Cell Simulation")
	defer rl.CloseWindow()

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
			cells[i].Move(rl.Vector2{X: 5, Y: 5})
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
