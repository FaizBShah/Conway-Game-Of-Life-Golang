package game

import (
	"errors"
	"reflect"

	"example.com/pkg/board"
	"example.com/pkg/cell"
	"example.com/pkg/utils"
)

type Game struct {
	board board.Board
}

func (g *Game) StartWithBoard(newBoard [][]cell.CellStatus) error {
	if !reflect.DeepEqual(g.board, board.Board{}) {
		return errors.New("game has already been started with a board")
	}

	if newBoard == nil {
		return errors.New("initialized with an empty board")
	}

	if len(newBoard) == 0 || len(newBoard[0]) == 0 {
		return errors.New("invalid board size")
	}

	newCreatedBoard, err := board.NewBoard(newBoard)

	if err != nil {
		return err
	}

	g.board = newCreatedBoard

	return nil
}

func (g *Game) StartWithRandomBoardOfSize(rows int, cols int) error {
	if !reflect.DeepEqual(g.board, board.Board{}) {
		return errors.New("game has already been started with a board")
	}

	if rows <= 0 || cols <= 0 {
		return errors.New("dimensions of the board should be valid")
	}

	starterBoard := make([][]cell.CellStatus, rows)

	for row := 0; row < rows; row++ {
		starterBoard[row] = make([]cell.CellStatus, cols)
		for col := 0; col < cols; col++ {
			starterBoard[row][col] = utils.RandomCellStatusSelector()
		}
	}

	newCreatedBoard, err := board.NewBoard(starterBoard)

	if err != nil {
		return err
	}

	g.board = newCreatedBoard

	return nil
}

func (g *Game) Tick() ([][]cell.CellStatus, error) {
	if reflect.DeepEqual(g.board, board.Board{}) {
		return nil, errors.New("game has not been initialized with a board yet")
	}

	return g.board.Next()
}

func (g *Game) TickN(n int) ([][]cell.CellStatus, error) {
	if reflect.DeepEqual(g.board, board.Board{}) {
		return nil, errors.New("game has not been initialized with a board yet")
	}

	for iter := 0; iter < n-1; iter++ {
		g.board.Next()
	}

	return g.board.Next()
}

func (g *Game) Stop() error {
	if reflect.DeepEqual(g.board, board.Board{}) {
		return errors.New("game has not been initialized with a board yet")
	}

	g.board = board.Board{}

	return nil
}
