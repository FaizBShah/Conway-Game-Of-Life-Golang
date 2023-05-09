package cell

import (
	"testing"
)

func TestShouldDeadCellIsAliveReturnFalse(t *testing.T) {
	cell := DeadCell{}

	if cell.IsAlive() {
		t.Errorf("Expected %v, but got %v", false, true)
	}
}

func TestShouldADeadCellBecomeAliveIfItHasExactlyThreeNeighbours(t *testing.T) {
	cell := DeadCell{}

	if res, err := cell.Next(3); err != nil || !res.IsAlive() {
		t.Errorf("Expected %v, but got %v", true, res.IsAlive())
	}
}

func TestShouldADeadCellRemainDeadIfItDoesNotHaveThreeNeighbours(t *testing.T) {
	cell1 := DeadCell{}
	cell2 := DeadCell{}

	if res, err := cell1.Next(2); err != nil || res.IsAlive() {
		t.Errorf("Expected %v, but got %v", false, res.IsAlive())
	}

	if res, err := cell2.Next(4); err != nil || res.IsAlive() {
		t.Errorf("Expected %v, but got %v", false, res.IsAlive())
	}
}

func TestShouldThrowErrorIfTryingToGetTheNextStateOfADeadCellWithLessThanZeroNeighbours(t *testing.T) {
	cell := DeadCell{}
	expectedErrMessage := "number of neighbours of a cell can only be between 0 and 8"

	if _, err := cell.Next(-1); err == nil || err.Error() != expectedErrMessage {
		t.Errorf("Expected an error object with message \"number of neighbours of a cell can only be between 0 and 8\"")
	}
}

func TestShouldThrowErrorIfTryingToGetTheNextStateOfAnDeadCellWithMoreThanEightNeighbours(t *testing.T) {
	cell := DeadCell{}
	expectedErrMessage := "number of neighbours of a cell can only be between 0 and 8"

	if _, err := cell.Next(9); err == nil || err.Error() != expectedErrMessage {
		t.Errorf("Expected an error object with message \"number of neighbours of a cell can only be between 0 and 8\"")
	}
}
