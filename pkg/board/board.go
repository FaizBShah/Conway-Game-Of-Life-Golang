package board

import (
	"errors"

	"example.com/pkg/cell"
	"example.com/pkg/utils"
)

type Board struct {
	board [][]cell.Cell
	rows  int
	cols  int
}

func (b *Board) Next() ([][]cell.CellStatus, error) {
	tempNewBoard := make([][]cell.Cell, b.rows)

	for row := 0; row < b.rows; row++ {
		tempNewBoard[row] = make([]cell.Cell, b.cols)
		for col := 0; col < b.cols; col++ {
			numberOfAliveNeighbours, err := utils.GetNumberOfAliveNeighbours(row, col, b.board)

			if err != nil {
				return nil, err
			}

			cell, err := b.board[row][col].Next(numberOfAliveNeighbours)

			if err != nil {
				return nil, err
			}

			tempNewBoard[row][col] = cell
		}
	}

	b.board = tempNewBoard

	return utils.ConvertCellBoardToCellStatusBoard(tempNewBoard)
}

func NewBoard(newBoard [][]cell.CellStatus) (Board, error) {
	if newBoard == nil {
		return Board{}, errors.New("initialized with an empty board")
	}

	if len(newBoard) == 0 || len(newBoard[0]) == 0 {
		return Board{}, errors.New("invalid board size")
	}

	rows := len(newBoard)
	cols := len(newBoard[0])

	board, err := utils.ConvertCellStatusBoardToCellBoard(newBoard)

	if err != nil {
		return Board{}, err
	}

	return Board{board, rows, cols}, nil
}
