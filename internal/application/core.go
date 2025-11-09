package application

import (
	"fmt"
	"os"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

func Run(input *InputParams, r Reader, w Writer, rand Random) error {
	switch input.cmd {
	case "generate":
		if err := generateMaze(input, w, rand); err != nil {
			return err
		}
	case "solve":
		if err := solveMaze(input, w, r); err != nil {
			return err
		}
	default:
		return fmt.Errorf("enter valid command: there are only generate and solve")
	}
	return nil
}

// Generator --------------------------------------------------------------------------------------------------
func generateMaze(input *InputParams, w Writer, rand Random) error {
	g := makeGenerator(input, rand)
	if g == nil {
		return fmt.Errorf("no such algorithm for generator: %s", input.algorithm)
	}
	maze := g.Generate(input.height, input.width)

	if err := w.PrintMaze(input.output, os.Stdout, maze); err != nil {
		return err
	}

	return nil
}

func makeGenerator(input *InputParams, rand Random) domain.Generator {
	switch input.algorithm {
	case "dfs":
		return &DFS{rnd: rand}
	case "prim":
		return &Prim{rnd: rand}
	default:
		return nil
	}
}

// Solver --------------------------------------------------------------------------------------------------
func solveMaze(input *InputParams, w Writer, r Reader) error {
	s := makeSolver(input)
	if s == nil {
		return fmt.Errorf("no such algorithm for solver: %s", input.algorithm)
	}

	maze, err := r.GetMazeFromFile(input.mazeFile)
	if err != nil {
		return err
	}

	width, height := (maze.GetWidth()-1)/2, (maze.GetHeight()-1)/2 // чтобы формат вывода точек
	// и размеров был ожидаем для пользователя, он же не знает о фокусах с удвоением размера лабиринта

	if !(input.startPoint.Row < 2*height+1 &&
		input.startPoint.Col < 2*width+1) {
		return fmt.Errorf("startPoint is out of bounds: %d,%d in shape (rows=%d, cols=%d) (0-index)",
			(input.startPoint.Row-1)/2, (input.startPoint.Col-1)/2, height, width) // аналогично
	}

	if !(input.endPoint.Row < 2*height+1 &&
		input.endPoint.Col < 2*width+1) {
		return fmt.Errorf("endPoint is out of bounds: %d,%d in shape (rows=%d, cols=%d) (0-index)",
			(input.endPoint.Row-1)/2, (input.endPoint.Col-1)/2, height, width) // аналогично
	}

	path := s.Solve(maze, input.startPoint, input.endPoint)

	if err := w.PrintPath(input.output, path, maze); err != nil {
		return err
	}

	return nil
}

func makeSolver(input *InputParams) domain.Solver {
	switch input.algorithm {
	case "astar":
		return &AStar{}
	case "dijkstra":
		return &Dijkstra{}
	default:
		return nil
	}
}
