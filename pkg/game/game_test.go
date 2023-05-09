package game

import (
	"reflect"
	"testing"

	"example.com/pkg/cell"
)

func TestShouldGameBeInitializedWithABoardSuccessfully(t *testing.T) {
	game := Game{}
	gameBoard := make([][]cell.CellStatus, 3)

	for i := 0; i < 3; i++ {
		gameBoard[i] = make([]cell.CellStatus, 3)
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

	if err := game.StartWithBoard(gameBoard); err != nil {
		t.Errorf("Expected not to throw an error")
	}

	if err := game.Stop(); err != nil {
		t.Errorf("Expected not to throw an error")
	}
}

func TestShouldThrowAnErrorIfTryingToInitializeAGameOnceItHasAlreadyStarted(t *testing.T) {
	game := Game{}
	gameBoard := make([][]cell.CellStatus, 3)

	for i := 0; i < 3; i++ {
		gameBoard[i] = make([]cell.CellStatus, 3)
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

	game.StartWithBoard(gameBoard)

	expectedErrMessage := "game has already been started with a board"

	if err := game.StartWithBoard(gameBoard); err == nil || err.Error() != expectedErrMessage {
		t.Errorf("Expected an error object with message \"game has already been started with a board")
	}
}

func TestShouldThrowAnErrorIfTryingToInitializeTheGameWithAnEmptyNullBoard(t *testing.T) {
	game := Game{}

	expectedErrMessage := "initialized with an empty board"

	if err := game.StartWithBoard(nil); err == nil || err.Error() != expectedErrMessage {
		t.Errorf("Expected an error object with message \"initialized with an empty board")
	}
}

func TestShouldThrowAnErrorIfTryingToInitializeTheGameWithInvalidBoardDimensions(t *testing.T) {
	game := Game{}
	gameBoard := make([][]cell.CellStatus, 0)

	expectedErrMessage := "invalid board size"

	if err := game.StartWithBoard(gameBoard); err == nil || err.Error() != expectedErrMessage {
		t.Errorf("Expected an error object with message \"invalid board size")
	}
}

func TestShouldGameBeInitializedWithARandomBoardSuccessfully(t *testing.T) {
	game := Game{}

	if err := game.StartWithRandomBoardOfSize(3, 3); err != nil {
		t.Errorf("Expected not to throw an error")
	}

	board, err := game.Tick()

	if err != nil {
		t.Errorf("Expected not to throw an error")
	}

	if len(board) != 3 || len(board[0]) != 3 {
		t.Errorf("Expected board of size 3 x 3")
	}

	if err := game.Stop(); err != nil {
		t.Errorf("Expected not to throw an error")
	}
}

func TestShouldThrowAnErrorIfTryingToRandomInitializeAGameOnceItHasAlreadyStarted(t *testing.T) {
	game := Game{}

	game.StartWithRandomBoardOfSize(3, 3)

	expectedErrMessage := "game has already been started with a board"

	if err := game.StartWithRandomBoardOfSize(3, 3); err == nil || err.Error() != expectedErrMessage {
		t.Errorf("Expected an error object with message \"game has already been started with a board")
	}
}

func TestShouldTThrowAnErrorIfTryingToRandomInitializeAGameWithNegativeRowOrColumn(t *testing.T) {
	game := Game{}

	expectedErrMessage := "dimensions of the board should be valid"

	if err := game.StartWithRandomBoardOfSize(-1, 3); err == nil || err.Error() != expectedErrMessage {
		t.Errorf("Expected an error object with message \"dimensions of the board should be valid")
	}

	if err := game.StartWithRandomBoardOfSize(3, -1); err == nil || err.Error() != expectedErrMessage {
		t.Errorf("Expected an error object with message \"dimensions of the board should be valid")
	}

	if err := game.StartWithRandomBoardOfSize(-1, -1); err == nil || err.Error() != expectedErrMessage {
		t.Errorf("Expected an error object with message \"dimensions of the board should be valid")
	}
}

func TestShouldTickWorkCorrectly(t *testing.T) {
	game := Game{}
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

	game.StartWithBoard(gameBoard)

	if result, err := game.Tick(); err != nil || !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestShouldThrowAnErrorIfTryingToCallTheTickFunctionWhenTheGameHasNotBeenInitializedWithABoard(t *testing.T) {
	game := Game{}

	expectedErrMessage := "game has not been initialized with a board yet"

	if _, err := game.Tick(); err == nil || err.Error() != expectedErrMessage {
		t.Errorf("Expected an error object with message \"game has not been initialized with a board yet")
	}
}

func TestShouldTickNTimesWorkCorrectly(t *testing.T) {
	game := Game{}
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
	expected[0][1] = cell.FILLED
	expected[0][2] = cell.EMPTY
	expected[1][0] = cell.FILLED
	expected[1][1] = cell.EMPTY
	expected[1][2] = cell.FILLED
	expected[2][0] = cell.FILLED
	expected[2][1] = cell.EMPTY
	expected[2][2] = cell.FILLED

	game.StartWithBoard(gameBoard)

	if result, err := game.TickN(2); err != nil || !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestShouldThrowAnErrorIfTryingToCallTheTickNFunctionWhenTheGameHasNotBeenInitializedWithABoard(t *testing.T) {
	game := Game{}

	expectedErrMessage := "game has not been initialized with a board yet"

	if _, err := game.TickN(5); err == nil || err.Error() != expectedErrMessage {
		t.Errorf("Expected an error object with message \"game has not been initialized with a board yet")
	}
}

func TestShouldStopFunctionWorkCorrectly(t *testing.T) {
	game := Game{}

	game.StartWithRandomBoardOfSize(3, 3)

	if err := game.Stop(); err != nil {
		t.Errorf("Expected not to throw an error")
	}
}

func TestShouldThrowAnErrorIfTryingToStopAGameWhichHasNotBeenInitialized(t *testing.T) {
	game := Game{}

	expectedErrMessage := "game has not been initialized with a board yet"

	if err := game.Stop(); err == nil || err.Error() != expectedErrMessage {
		t.Errorf("Expected an error object with message \"game has not been initialized with a board yet")
	}
}
