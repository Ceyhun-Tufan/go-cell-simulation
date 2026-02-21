package main

type blueCells struct {
	pos   Vector2
	speed Vector2
}

func (cell *blueCells) Move(speed Vector2) {
	cell.pos.x += speed.x
	cell.pos.y += speed.y
}
