package cell

type AliveCell struct{}

func (a AliveCell) IsAlive() bool {
	return true
}

func (a AliveCell) Next(numberOfAliveNeighbours int) (Cell, error) {
	return nextCell(a.IsAlive(), numberOfAliveNeighbours)
}
