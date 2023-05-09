package utils

import (
	"errors"

	"example.com/pkg/cell"
)

func ConvertCellBoardToCellStatusBoard(board [][]cell.Cell) ([][]cell.CellStatus, error) {
	if board == nil {
		return nil, errors.New("initialized with an empty board")
	}

	if len(board) == 0 || len(board[0]) == 0 {
		return nil, errors.New("invalid board size")
	}

	rows := len(board)
	cols := len(board[0])

	resBoard := make([][]cell.CellStatus, rows)

	for row := 0; row < rows; row++ {
		resBoard[row] = make([]cell.CellStatus, cols)
		for col := 0; col < cols; col++ {
			if board[row][col].IsAlive() {
				resBoard[row][col] = cell.FILLED
			} else {
				resBoard[row][col] = cell.EMPTY
			}
		}
	}

	return resBoard, nil
}

func ConvertCellStatusBoardToCellBoard(board [][]cell.CellStatus) ([][]cell.Cell, error) {
	if board == nil {
		return nil, errors.New("initialized with an empty board")
	}

	if len(board) == 0 || len(board[0]) == 0 {
		return nil, errors.New("invalid board size")
	}

	rows := len(board)
	cols := len(board[0])

	resBoard := make([][]cell.Cell, rows)

	for row := 0; row < rows; row++ {
		resBoard[row] = make([]cell.Cell, cols)
		for col := 0; col < cols; col++ {
			if board[row][col] == cell.FILLED {
				resBoard[row][col] = cell.AliveCell{}
			} else {
				resBoard[row][col] = cell.DeadCell{}
			}
		}
	}

	return resBoard, nil
}
