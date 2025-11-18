package application

import (
	"io"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

type DFS struct {
	rnd Random
}

type Prim struct {
	rnd Random
}

type Dijkstra struct {
}

type AStar struct {
}

type Reader interface {
	// Maze -----------------------------------------------
	GetMazeFromFile(path string) (*domain.Maze, error)
}

type Writer interface {
	PrintMaze(path string, buf io.Writer, maze *domain.Maze) error
	PrintPath(output string, path *domain.Path, maze *domain.Maze) error
}

type Random interface {
	GetRandom(num int) int
}

type pqItem struct {
	dist int
	c    domain.Cell
}

type astarItem struct {
	dist int
	h    int
	c    domain.Cell
}

type PQDijkstra []pqItem

type PQAStar []astarItem
