package utils

import (
	"testing"

	"example.com/pkg/cell"
)

func TestShouldFindTheCorrectNumberOfAliveNeighboursOfACell(t *testing.T) {
	gameBoard := make([][]cell.Cell, 3)

	for i := 0; i < 3; i++ {
		gameBoard[i] = make([]cell.Cell, 3)
	}

	gameBoard[0][0] = cell.DeadCell{}
	gameBoard[0][1] = cell.DeadCell{}
	gameBoard[0][2] = cell.AliveCell{}
	gameBoard[1][0] = cell.DeadCell{}
	gameBoard[1][1] = cell.AliveCell{}
	gameBoard[1][2] = cell.DeadCell{}
	gameBoard[2][0] = cell.AliveCell{}
	gameBoard[2][1] = cell.AliveCell{}
	gameBoard[2][2] = cell.DeadCell{}

	if result, err := GetNumberOfAliveNeighbours(1, 1, gameBoard); err != nil || result != 3 {
		t.Errorf("Expected %v, but got %v", 3, result)
	}
}

func TestShouldThrowAnErrorIfTryingToPassANullBoardObjectDuringFindingTheNumberOfAliveNeighbours(t *testing.T) {
	expectedErrMessage := "initialized with an empty board"

	if _, err := GetNumberOfAliveNeighbours(1, 1, nil); err == nil || err.Error() != expectedErrMessage {
		t.Errorf("Expected an error object with message \"initialized with an empty board")
	}
}

func TestShouldThrowAnErrorIfTryingToPassABoardOfInvalidLength(t *testing.T) {
	gameBoard := make([][]cell.Cell, 0)

	expectedErrMessage := "invalid board size"

	if _, err := GetNumberOfAliveNeighbours(1, 1, gameBoard); err == nil || err.Error() != expectedErrMessage {
		t.Errorf("Expected an error object with message \"invalid board size")
	}
}

func TestShouldThrowAnErrorIfTryingToPassInvalidCellCoordinates(t *testing.T) {
	gameBoard := make([][]cell.Cell, 3)

	for i := 0; i < 3; i++ {
		gameBoard[i] = make([]cell.Cell, 3)
	}

	gameBoard[0][0] = cell.DeadCell{}
	gameBoard[0][1] = cell.DeadCell{}
	gameBoard[0][2] = cell.AliveCell{}
	gameBoard[1][0] = cell.DeadCell{}
	gameBoard[1][1] = cell.AliveCell{}
	gameBoard[1][2] = cell.DeadCell{}
	gameBoard[2][0] = cell.AliveCell{}
	gameBoard[2][1] = cell.AliveCell{}
	gameBoard[2][2] = cell.DeadCell{}

	expectedErrMessage := "coordinates of the cell are out of bounds of the board size"

	if _, err := GetNumberOfAliveNeighbours(4, -1, gameBoard); err == nil || err.Error() != expectedErrMessage {
		t.Errorf("Expected an error object with message \"coordinates of the cell are out of bounds of the board size")
	}
}
