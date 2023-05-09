package utils

import (
	"math/rand"
	"time"

	"example.com/pkg/cell"
)

func RandomCellStatusSelector() cell.CellStatus {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	cellStatuses := [...]cell.CellStatus{cell.FILLED, cell.EMPTY}
	n := len(cellStatuses)
	return cellStatuses[r.Intn(n)]
}
