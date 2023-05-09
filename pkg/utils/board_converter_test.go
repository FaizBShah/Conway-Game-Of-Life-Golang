package utils

import (
	"reflect"
	"testing"

	"example.com/pkg/cell"
)

func TestShouldConvertCellBoardToCellStatusBoardCorrectly(t *testing.T) {
	gameBoard := make([][]cell.Cell, 3)
	expected := make([][]cell.CellStatus, 3)

	for i := 0; i < 3; i++ {
		gameBoard[i] = make([]cell.Cell, 3)
		expected[i] = make([]cell.CellStatus, 3)
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

	expected[0][0] = cell.EMPTY
	expected[0][1] = cell.EMPTY
	expected[0][2] = cell.FILLED
	expected[1][0] = cell.EMPTY
	expected[1][1] = cell.FILLED
	expected[1][2] = cell.EMPTY
	expected[2][0] = cell.FILLED
	expected[2][1] = cell.FILLED
	expected[2][2] = cell.EMPTY

	if result, err := ConvertCellBoardToCellStatusBoard(gameBoard); err != nil || !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestShouldThrowAnErrorIfTryingToPassANullObjectDuringCellBoardToCellStatusBoardConversion(t *testing.T) {
	expectedErrMessage := "initialized with an empty board"

	if _, err := ConvertCellBoardToCellStatusBoard(nil); err == nil || err.Error() != expectedErrMessage {
		t.Errorf("Expected an error object with message \"initialized with an empty board")
	}
}

func TestShouldThrowAnErrorIfTryingToPassACellStatusBoardOfInvalidLengthDuringCellStatusBoardToCellBoardConversion(t *testing.T) {
	gameBoard := make([][]cell.Cell, 0)

	expectedErrMessage := "invalid board size"

	if _, err := ConvertCellBoardToCellStatusBoard(gameBoard); err == nil || err.Error() != expectedErrMessage {
		t.Errorf("Expected an error object with message \"invalid board size")
	}
}

func TestShouldConvertCellStatusBoardToCellBoardCorrectly(t *testing.T) {
	gameBoard := make([][]cell.CellStatus, 3)
	expected := make([][]cell.Cell, 3)

	for i := 0; i < 3; i++ {
		gameBoard[i] = make([]cell.CellStatus, 3)
		expected[i] = make([]cell.Cell, 3)
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

	expected[0][0] = cell.DeadCell{}
	expected[0][1] = cell.DeadCell{}
	expected[0][2] = cell.AliveCell{}
	expected[1][0] = cell.DeadCell{}
	expected[1][1] = cell.AliveCell{}
	expected[1][2] = cell.DeadCell{}
	expected[2][0] = cell.AliveCell{}
	expected[2][1] = cell.AliveCell{}
	expected[2][2] = cell.DeadCell{}

	if result, err := ConvertCellStatusBoardToCellBoard(gameBoard); err != nil || !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestShouldThrowAnErrorIfTryingToPassANullObjectDuringCellStatusBoardToCellBoardConversion(t *testing.T) {
	expectedErrMessage := "initialized with an empty board"

	if _, err := ConvertCellStatusBoardToCellBoard(nil); err == nil || err.Error() != expectedErrMessage {
		t.Errorf("Expected an error object with message \"initialized with an empty board")
	}
}

func TestShouldThrowAnErrorIfTryingToPassACellBoardOfInvalidLengthDuringCellStatusBoardToCellStatusBoardConversion(t *testing.T) {
	gameBoard := make([][]cell.CellStatus, 0)

	expectedErrMessage := "invalid board size"

	if _, err := ConvertCellStatusBoardToCellBoard(gameBoard); err == nil || err.Error() != expectedErrMessage {
		t.Errorf("Expected an error object with message \"invalid board size")
	}
}
