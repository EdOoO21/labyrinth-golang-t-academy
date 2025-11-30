package support

import (
	cells "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/cell"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Heuristic(c, end cells.Cell) int {
	return Abs(c.Row-end.Row) + Abs(c.Col-end.Col)
}

func CheckPath(scheme [][]cells.CellType, maze *maze.Maze, path []cells.Cell) bool {
	var prev cells.Cell
	if len(path) == 0 {
		return true
	}

	for i, p := range path {
		cell := cells.CellPath
		switch i {
		case 0:
			cell = cells.CellStart
		case len(path) - 1:
			cell = cells.CellEnd
		}

		c, _ := maze.GetCell(p.Row, p.Col)
		if !((scheme[p.Row][p.Col] == cells.CellEmpty) && (c == cell)) {
			return false
		}

		if i != 0 && (Abs(prev.Row-p.Row)+Abs(prev.Col-p.Col)) != 1 {
			return false
		}
		prev = p
	}

	return true
}
