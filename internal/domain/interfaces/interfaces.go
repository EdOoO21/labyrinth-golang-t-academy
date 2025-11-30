package interfaces

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/cell"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/path"
)

type Generator interface {
	Generate(rows, cols int) *maze.Maze
}

type Solver interface {
	Solve(maze *maze.Maze, start, end cell.Cell) *path.Path
}
