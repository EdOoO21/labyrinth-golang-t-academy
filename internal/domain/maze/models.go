package maze

import cell "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/cell"

type Maze struct {
	scheme        [][]cell.CellType
	width, height int
}
