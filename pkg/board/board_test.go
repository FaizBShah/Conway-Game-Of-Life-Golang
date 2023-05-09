package board

import (
	"reflect"
	"testing"

	"example.com/pkg/cell"
)

func TestShouldThrowAnErrorIfTryingToPassANullObjectDuringBoardObjectCreation(t *testing.T) {
	expectedErrMessage := "initialized with an empty board"

	if _, err := NewBoard(nil); err == nil || err.Error() != expectedErrMessage {
		t.Errorf("Expected an error object with message \"initialized with an empty board")
	}
}

func TestShouldThrowAnErrorIfTryingToPassABoardOfInvalidLength(t *testing.T) {
	gameBoard := make([][]cell.CellStatus, 0)

	expectedErrMessage := "invalid board size"

	if _, err := NewBoard(gameBoard); err == nil || err.Error() != expectedErrMessage {
		t.Errorf("Expected an error object with message \"invalid board size")
	}
}

func TestShouldGenerateCorrectNextBoard(t *testing.T) {
	gameBoard := make([][]cell.CellStatus, 3)
	expected := make([][]cell.CellStatus, 3)

	for i := 0; i < 3; i++ {
		gameBoard[i] = make([]cell.CellStatus, 3)
		expected[i] = make([]cell.CellStatus, 3)
	}

	gameBoard[0][0] = cell.EMPTY
	gameBoard[0][1] = cell.EMPTY
	gameBoard[0][2] = cell.FILLED
	gameBoard[1][0] = cell.EMPTY
	gameBoard[1][1] = cell.FILLED
	gameBoard[1][2] = cell.EMPTY
	gameBoard[2][0] = cell.FILLED
	gameBoard[2][1] = cell.FILLED
	gameBoard[2][2] = cell.EMPTY

	expected[0][0] = cell.EMPTY
	expected[0][1] = cell.EMPTY
	expected[0][2] = cell.EMPTY
	expected[1][0] = cell.FILLED
	expected[1][1] = cell.FILLED
	expected[1][2] = cell.FILLED
	expected[2][0] = cell.FILLED
	expected[2][1] = cell.FILLED
	expected[2][2] = cell.EMPTY

	board, err := NewBoard(gameBoard)

	if err != nil {
		t.Errorf(err.Error())
	}

	if result, err := board.Next(); err != nil || !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
