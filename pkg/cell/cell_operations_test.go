package cell

import (
	"reflect"
	"testing"
)

func TestShouldReturnDeadCellIfNumberOfAliveNeighboursIsZero(t *testing.T) {
	if cell, err := nextCell(true, 0); err != nil || !reflect.DeepEqual(DeadCell{}, cell) {
		t.Errorf("Expected %v, but got %v", DeadCell{}, cell)
	}

	if cell, err := nextCell(false, 0); err != nil || !reflect.DeepEqual(DeadCell{}, cell) {
		t.Errorf("Expected %v, but got %v", DeadCell{}, cell)
	}
}

func TestShouldReturnDeadCellIfNumberOfAliveNeighboursIsOne(t *testing.T) {
	if cell, err := nextCell(true, 0); err != nil || !reflect.DeepEqual(DeadCell{}, cell) {
		t.Errorf("Expected %v, but got %v", DeadCell{}, cell)
	}

	if cell, err := nextCell(false, 0); err != nil || !reflect.DeepEqual(DeadCell{}, cell) {
		t.Errorf("Expected %v, but got %v", DeadCell{}, cell)
	}
}

func TestShouldReturnAliveCellIfNumberOfNeighboursIsTwoAndCellIsAlive(t *testing.T) {
	if cell, err := nextCell(true, 2); err != nil || !reflect.DeepEqual(AliveCell{}, cell) {
		t.Errorf("Expected %v, but got %v", AliveCell{}, cell)
	}
}

func TestShouldReturnDeadCellIfNumberOfNeighboursIsTwoAndCellIsDead(t *testing.T) {
	if cell, err := nextCell(false, 2); err != nil || !reflect.DeepEqual(DeadCell{}, cell) {
		t.Errorf("Expected %v, but got %v", DeadCell{}, cell)
	}
}

func TestShouldReturnAliveCellIfNumberOfNeighboursIsThreeAndCellIsAlive(t *testing.T) {
	if cell, err := nextCell(true, 3); err != nil || !reflect.DeepEqual(AliveCell{}, cell) {
		t.Errorf("Expected %v, but got %v", AliveCell{}, cell)
	}
}

func TestShouldReturnAliveCellIfNumberOfNeighboursIsThreeAndCellIsDead(t *testing.T) {
	if cell, err := nextCell(false, 3); err != nil || !reflect.DeepEqual(AliveCell{}, cell) {
		t.Errorf("Expected %v, but got %v", AliveCell{}, cell)
	}
}

func TestShouldReturnDeadCellIfNumberOfAliveNeighboursIsFour(t *testing.T) {
	if cell, err := nextCell(true, 4); err != nil || !reflect.DeepEqual(DeadCell{}, cell) {
		t.Errorf("Expected %v, but got %v", DeadCell{}, cell)
	}

	if cell, err := nextCell(false, 4); err != nil || !reflect.DeepEqual(DeadCell{}, cell) {
		t.Errorf("Expected %v, but got %v", DeadCell{}, cell)
	}
}

func TestShouldReturnDeadCellIfNumberOfAliveNeighboursIsFive(t *testing.T) {
	if cell, err := nextCell(true, 5); err != nil || !reflect.DeepEqual(DeadCell{}, cell) {
		t.Errorf("Expected %v, but got %v", DeadCell{}, cell)
	}

	if cell, err := nextCell(false, 5); err != nil || !reflect.DeepEqual(DeadCell{}, cell) {
		t.Errorf("Expected %v, but got %v", DeadCell{}, cell)
	}
}

func TestShouldReturnDeadCellIfNumberOfAliveNeighboursIsSix(t *testing.T) {
	if cell, err := nextCell(true, 6); err != nil || !reflect.DeepEqual(DeadCell{}, cell) {
		t.Errorf("Expected %v, but got %v", DeadCell{}, cell)
	}

	if cell, err := nextCell(false, 6); err != nil || !reflect.DeepEqual(DeadCell{}, cell) {
		t.Errorf("Expected %v, but got %v", DeadCell{}, cell)
	}
}

func TestShouldReturnDeadCellIfNumberOfAliveNeighboursIsSeven(t *testing.T) {
	if cell, err := nextCell(true, 7); err != nil || !reflect.DeepEqual(DeadCell{}, cell) {
		t.Errorf("Expected %v, but got %v", DeadCell{}, cell)
	}

	if cell, err := nextCell(false, 7); err != nil || !reflect.DeepEqual(DeadCell{}, cell) {
		t.Errorf("Expected %v, but got %v", DeadCell{}, cell)
	}
}

func TestShouldReturnDeadCellIfNumberOfAliveNeighboursIsEight(t *testing.T) {
	if cell, err := nextCell(true, 8); err != nil || !reflect.DeepEqual(DeadCell{}, cell) {
		t.Errorf("Expected %v, but got %v", DeadCell{}, cell)
	}

	if cell, err := nextCell(false, 8); err != nil || !reflect.DeepEqual(DeadCell{}, cell) {
		t.Errorf("Expected %v, but got %v", DeadCell{}, cell)
	}
}

func TestShouldThrowAnErrorIfNumberOfAliveNeighboursIsLessThanZero(t *testing.T) {
	expectedErrMessage := "number of neighbours of a cell can only be between 0 and 8"

	if _, err := nextCell(true, -1); err == nil || err.Error() != expectedErrMessage {
		t.Errorf("Expected an error object with message \"number of neighbours of a cell can only be between 0 and 8\"")
	}

	if _, err := nextCell(false, -1); err == nil || err.Error() != expectedErrMessage {
		t.Errorf("Expected an error object with message \"number of neighbours of a cell can only be between 0 and 8\"")
	}
}

func TestShouldThrowAnErrorIfNumberOfAliveNeighboursIsMoreThanEight(t *testing.T) {
	expectedErrMessage := "number of neighbours of a cell can only be between 0 and 8"

	if _, err := nextCell(true, 9); err == nil || err.Error() != expectedErrMessage {
		t.Errorf("Expected an error object with message \"number of neighbours of a cell can only be between 0 and 8\"")
	}

	if _, err := nextCell(false, 9); err == nil || err.Error() != expectedErrMessage {
		t.Errorf("Expected an error object with message \"number of neighbours of a cell can only be between 0 and 8\"")
	}
}
