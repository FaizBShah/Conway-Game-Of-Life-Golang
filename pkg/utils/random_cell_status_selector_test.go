package utils

import (
	"testing"

	"example.com/pkg/cell"
)

func TestShouldRandomCellStatusSelectorWorkCorrectly(t *testing.T) {
	for i := 0; i < 1000; i++ {
		if cellStatus := RandomCellStatusSelector(); cellStatus != cell.FILLED && cellStatus != cell.EMPTY {
			t.Errorf("Expected random function to return either EMPTY or FILLED")
		}
	}
}
