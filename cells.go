package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Cell struct {
	pos              rl.Vector2
	speed            rl.Vector2
	radius           float32
	densityThreshold int
	separation       int
	color            rl.Color
}

// default needs

func NewBlueCell() Cell {
	return Cell{
		pos:              rl.Vector2{X: rand.Float32() * float32(SCREEN_WIDTH), Y: rand.Float32() * float32(SCREEN_HEIGHT)},
		speed:            rl.Vector2{X: 0, Y: 0},
		radius:           5,
		densityThreshold: 10,
		color:            rl.Blue,
	}
}

func NewRedCell() Cell {
	return Cell{
		pos:              rl.Vector2{X: rand.Float32() * float32(SCREEN_WIDTH), Y: rand.Float32() * float32(SCREEN_HEIGHT)},
		speed:            rl.Vector2{X: 0, Y: 0},
		radius:           3,
		densityThreshold: 5,
		color:            rl.Red,
	}
}

func (cell *Cell) moveTowardsAttraction() { // no need for vector2 here

	tempx := cell.pos.X + cell.speed.X
	tempy := cell.pos.Y + cell.speed.Y

	if tempx < 0 || tempx > float32(SCREEN_WIDTH) {
		return
	}
	cell.pos.X = tempx
	if tempy < 0 || tempy > float32(SCREEN_HEIGHT) {
		return
	}

	cell.pos.Y = tempy

}

// to find the WAY it wants to go
func (cell *Cell) findTheAttractionVector(neighbors []int) {
	// neighbors is a 1D array that is originally a 3x3 grid

	force := rl.Vector2{X: 0, Y: 0}

	for i := range neighbors {
		row := i/3 - 1
		col := i%3 - 1
		if neighbors[i] < cell.densityThreshold {
			force.X += float32(col * neighbors[i])
			force.Y += float32(row * neighbors[i])
		} else {
			force.X -= float32(col * neighbors[i])
			force.Y -= float32(row * neighbors[i])
		}

	}

	cell.speed = rl.Vector2Add(cell.speed, force)
	if cell.speed.X > MAX_CELL_SPEED {
		cell.speed.X = MAX_CELL_SPEED
	}
	if cell.speed.X < MAX_CELL_SPEED*-1 {
		cell.speed.X = MAX_CELL_SPEED * -1
	}
	if cell.speed.Y > MAX_CELL_SPEED {
		cell.speed.Y = MAX_CELL_SPEED
	}
	if cell.speed.Y < MAX_CELL_SPEED*-1 {
		cell.speed.Y = MAX_CELL_SPEED * -1
	}
	// cell.speed = rl.Vector2Normalize(cell.speed)
}
