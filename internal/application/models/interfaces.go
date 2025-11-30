package models

import (
	"io"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/path"
)

type Reader interface {
	// Maze -----------------------------------------------
	GetMazeFromFile(path string) (*maze.Maze, error)
}

type Writer interface {
	PrintMaze(path string, buf io.Writer, maze *maze.Maze) error
	PrintPath(output string, path *path.Path, maze *maze.Maze) error
}

type Random interface {
	GetRandom(num int) int
}
