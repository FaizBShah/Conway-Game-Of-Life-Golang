package cell

import (
	"testing"
)

func TestShouldAliveCellIsAliveReturnTrue(t *testing.T) {
	cell := AliveCell{}

	if !cell.IsAlive() {
		t.Errorf("Expected %v, but got %v", true, false)
	}
}

func TestShouldAnAliveCellDieIfItHasLessThanTwoNeighbours(t *testing.T) {
	cell1 := AliveCell{}
	cell2 := AliveCell{}

	if res, err := cell1.Next(0); err != nil || res.IsAlive() {
		t.Errorf("Expected %v, but got %v", false, res.IsAlive())
	}

	if res, err := cell2.Next(1); err != nil || res.IsAlive() {
		t.Errorf("Expected %v, but got %v", false, res.IsAlive())
	}
}

func TestShouldAnAliveCellRemainAliveIfItHasTwoOrThreeNeighbours(t *testing.T) {
	cell1 := AliveCell{}
	cell2 := AliveCell{}

	if res, err := cell1.Next(2); err != nil || !res.IsAlive() {
		t.Errorf("Expected %v, but got %v", true, res.IsAlive())
	}

	if res, err := cell2.Next(3); err != nil || !res.IsAlive() {
		t.Errorf("Expected %v, but got %v", true, res.IsAlive())
	}
}

func TestShouldAnAliveCellDieIfItHasMoreThanThreeNeighbours(t *testing.T) {
	cell1 := AliveCell{}
	cell2 := AliveCell{}

	if res, err := cell1.Next(4); err != nil || res.IsAlive() {
		t.Errorf("Expected %v, but got %v", false, res.IsAlive())
	}

	if res, err := cell2.Next(7); err != nil || res.IsAlive() {
		t.Errorf("Expected %v, but got %v", false, res.IsAlive())
	}
}

func TestShouldThrowErrorIfTryingToGetTheNextStateOfAnAliveCellWithLessThanZeroNeighbours(t *testing.T) {
	cell := AliveCell{}
	expectedErrMessage := "number of neighbours of a cell can only be between 0 and 8"

	if _, err := cell.Next(-1); err == nil || err.Error() != expectedErrMessage {
		t.Errorf("Expected an error object with message \"number of neighbours of a cell can only be between 0 and 8\"")
	}
}

func TestShouldThrowErrorIfTryingToGetTheNextStateOfAnAliveCellWithMoreThanEightNeighbours(t *testing.T) {
	cell := AliveCell{}
	expectedErrMessage := "number of neighbours of a cell can only be between 0 and 8"

	if _, err := cell.Next(9); err == nil || err.Error() != expectedErrMessage {
		t.Errorf("Expected an error object with message \"number of neighbours of a cell can only be between 0 and 8\"")
	}
}
