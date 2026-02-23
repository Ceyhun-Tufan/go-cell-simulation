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
	color            rl.Color
}

// default needs

func NewBlueCell() Cell {
	return Cell{
		pos:              rl.Vector2{X: rand.Float32() * float32(SCREEN_WIDTH), Y: rand.Float32() * float32(SCREEN_HEIGHT)},
		speed:            rl.Vector2{X: 0, Y: 0},
		radius:           4,
		densityThreshold: 15,
		color:            rl.Blue,
	}
}

func (cell *Cell) MoveTowardsAttraction(speed rl.Vector2) { // no need for vector2 here
	tempx := cell.pos.X + speed.X
	tempy := cell.pos.Y + speed.Y

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
// func (cell *Cell) findTheAttractionVector(neighbors []int) {

// 	attVector := rl.Vector2{X: 0, Y: 0}

// 	for i := range neighbors {

// 	}

// }
