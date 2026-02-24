package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// consts

const SCREEN_HEIGHT int32 = 600
const SCREEN_WIDTH int32 = 800

const CELL_COUNT int = 2000
const GRIDCELL_SIZE int = 20

const MAX_CELL_SPEED float32 = 1

var pause bool = false

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

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		if rl.IsKeyPressed(rl.KeySpace) {
			pause = !pause
		}

		// if pause {

		// }

		grid.clear()
		// fmt.Println("cleared")
		// fmt.Printf("len(grid.gridCells): %v\n", len(grid.gridCells))
		grid.insertCellsToGrid(cells)
		// fmt.Println("inserted")
		// fmt.Printf("len(grid.gridCells): %v\n", len(grid.gridCells))

		// draw
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		for i := range cells {
			tmp := cells[i]
			cells[i].findTheAttractionVector(grid.findNeighbors(i, cells))
			cells[i].moveTowardsAttraction() // need to normalize movement later on
			rl.DrawCircleV(tmp.pos, tmp.radius, tmp.color)
		}

		rl.EndDrawing()
	}
}

func initCells() {
	for range CELL_COUNT / 2 {
		cells = append(cells, NewBlueCell())
	}
	for range CELL_COUNT / 2 {
		cells = append(cells, NewRedCell())
	}
}
