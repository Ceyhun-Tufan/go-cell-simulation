package main

type Grid struct {
	cellSize  int
	cols      int
	rows      int
	gridCells [][]int // indisleri tutmak icin
}

func NewGrid(width int, height int, gridcellSize int) *Grid {
	// be careful that cellsize being an integer

	cols := width / gridcellSize
	rows := height / gridcellSize

	gridCells := make([][]int, cols*rows)

	return &Grid{
		cellSize:  gridcellSize,
		cols:      cols,
		rows:      rows,
		gridCells: gridCells,
	}

}

func (g *Grid) gridCellIndex(x, y float32) int {

	var col int = int(x) / g.cellSize
	var row int = int(y) / g.cellSize

	if col < 0 {
		col = 0
	}
	if row < 0 {
		row = 0
	}
	if col > g.cols {
		col = g.cols - 1
	}
	if row > g.rows {
		row = g.rows - 1
	}

	return row*g.cols + col

}

func (g *Grid) clear() {
	for i := range g.gridCells {
		g.gridCells[i] = g.gridCells[i][:0] // instead of setting the array from scratch, just zeroing it
	}

}

func (g *Grid) insertCellsToGrid(cells []Cell) {

	// finding the cells index
	// then appending it to the gridCell[idx] array - agents idx are stored there

	for i := range cells {
		idx := g.gridCellIndex(cells[i].pos.X, cells[i].pos.Y)
		g.gridCells[idx] = append(g.gridCells[idx], i)
	}
}

func (g *Grid) findNeighbors(cellIdx int, cells []Cell) []int {
	// to make a real cell like structers, i need to find a way to attract them together
	// 1D slice makes thing a little harder to understand sometimes
	// nrow*g.cols + ncol => shit ass technique
	cell := cells[cellIdx]

	var col int = int(cell.pos.X) / g.cellSize
	var row int = int(cell.pos.Y) / g.cellSize

	var neighborsCount []int

	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			// n = neighbor
			ncol := col + dx
			nrow := row + dy

			if ncol < 0 || nrow < 0 || ncol >= g.cols || nrow >= g.rows {
				neighborsCount = append(neighborsCount, 0)
				continue
			}
			// ncIdx is the index of the neighbor cell in the gridCells
			var ncIdx int = nrow*g.cols + ncol
			neighborsCount = append(neighborsCount, len(g.gridCells[ncIdx]))

		}
	}

	return neighborsCount

}
