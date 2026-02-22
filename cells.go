package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Cell struct {
	pos     rl.Vector2
	speed   rl.Vector2
	radius  float32
	kind    int
	attract int
	color   rl.Color
}

// default needs

func NewBlueCell() Cell {
	return Cell{
		pos:    rl.Vector2{X: rand.Float32() * float32(SCREEN_WIDTH), Y: rand.Float32() * float32(SCREEN_HEIGHT)},
		speed:  rl.Vector2{X: 0, Y: 0},
		radius: 4,
		color:  rl.Blue,
	}
}

func (cell *Cell) Move(speed rl.Vector2) {
	cell.pos.X += speed.X
	cell.pos.X += speed.Y
}
