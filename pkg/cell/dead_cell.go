package cell

type DeadCell struct{}

func (d DeadCell) IsAlive() bool {
	return false
}

func (d DeadCell) Next(numberOfAliveNeighbours int) (Cell, error) {
	return nextCell(d.IsAlive(), numberOfAliveNeighbours)
}
