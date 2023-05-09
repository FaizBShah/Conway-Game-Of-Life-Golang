package utils

import (
	"errors"

	"example.com/pkg/cell"
)

func GetNumberOfAliveNeighbours(row int, col int, board [][]cell.Cell) (int, error) {
	if board == nil {
		return 0, errors.New("initialized with an empty board")
	}

	if len(board) == 0 || len(board[0]) == 0 {
		return 0, errors.New("invalid board size")
	}

	if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) {
		return 0, errors.New("coordinates of the cell are out of bounds of the board size")
	}

	xCoordinatesForNeighbours := [8]int{-1, 0, 1, -1, 1, -1, 0, 1}
	yCoordinatesForNeighbours := [8]int{1, 1, 1, 0, 0, -1, -1, -1}

	count := 0

	for k := 0; k < 8; k++ {
		xPos := row + xCoordinatesForNeighbours[k]
		yPos := col + yCoordinatesForNeighbours[k]

		if xPos < 0 || yPos < 0 || xPos >= len(board) || yPos >= len(board[0]) {
			continue
		}

		if board[xPos][yPos].IsAlive() {
			count++
		}
	}

	return count, nil
}
