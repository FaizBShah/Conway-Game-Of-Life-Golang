package cell

import "errors"

var cellFunctionsMapping = map[int]func(bool) bool{
	2: func(isAlive bool) bool { return isAlive },
	3: func(isAlive bool) bool { return true },
}

func getCellFunction(num int) func(bool) bool {
	if v, ok := cellFunctionsMapping[num]; ok {
		return v
	}

	return func(b bool) bool { return false }
}

func nextCell(isAlive bool, numberOfAliveNeighbours int) (Cell, error) {
	if numberOfAliveNeighbours < 0 || numberOfAliveNeighbours > 8 {
		return nil, errors.New("number of neighbours of a cell can only be between 0 and 8")
	}

	cellFunction := getCellFunction(numberOfAliveNeighbours)

	if cellFunction(isAlive) {
		return AliveCell{}, nil
	}

	return DeadCell{}, nil
}
