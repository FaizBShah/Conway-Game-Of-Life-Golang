package cell

type Cell interface {
	IsAlive() bool
	Next(numberOfAliveNeighbours int) (Cell, error)
}
