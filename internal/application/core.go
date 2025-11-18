package application

import (
	"fmt"
	"os"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/parsers"
)

func Run(input *parsers.InputParams, r Reader, w Writer, rand Random) error {
	switch input.Cmd {
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
func generateMaze(input *parsers.InputParams, w Writer, rand Random) error {
	g := makeGenerator(input, rand)
	if g == nil {
		return fmt.Errorf("no such algorithm for generator: %s", input.Algorithm)
	}
	maze := g.Generate(input.Height, input.Width)

	if err := w.PrintMaze(input.Output, os.Stdout, maze); err != nil {
		return err
	}

	return nil
}

func makeGenerator(input *parsers.InputParams, rand Random) domain.Generator {
	switch input.Algorithm {
	case "dfs":
		return &DFS{rnd: rand}
	case "prim":
		return &Prim{rnd: rand}
	default:
		return nil
	}
}

// Solver --------------------------------------------------------------------------------------------------
func solveMaze(input *parsers.InputParams, w Writer, r Reader) error {
	s := makeSolver(input)
	if s == nil {
		return fmt.Errorf("no such algorithm for solver: %s", input.Algorithm)
	}

	maze, err := r.GetMazeFromFile(input.MazeFile)
	if err != nil {
		return err
	}

	width, height := (maze.GetWidth()-1)/2, (maze.GetHeight()-1)/2 // чтобы формат вывода точек
	// и размеров был ожидаем для пользователя, он же не знает о фокусах с удвоением размера лабиринта

	if !(input.StartPoint.Row < 2*height+1 &&
		input.StartPoint.Col < 2*width+1) {
		return fmt.Errorf("startPoint is out of bounds: %d,%d in shape (rows=%d, cols=%d) (0-index)",
			(input.StartPoint.Row-1)/2, (input.StartPoint.Col-1)/2, height, width) // аналогично
	}

	if !(input.EndPoint.Row < 2*height+1 &&
		input.EndPoint.Col < 2*width+1) {
		return fmt.Errorf("endPoint is out of bounds: %d,%d in shape (rows=%d, cols=%d) (0-index)",
			(input.EndPoint.Row-1)/2, (input.EndPoint.Col-1)/2, height, width) // аналогично
	}

	path := s.Solve(maze, input.StartPoint, input.EndPoint)

	if err := w.PrintPath(input.Output, path, maze); err != nil {
		return err
	}

	return nil
}

func makeSolver(input *parsers.InputParams) domain.Solver {
	switch input.Algorithm {
	case "astar":
		return &AStar{}
	case "dijkstra":
		return &Dijkstra{}
	default:
		return nil
	}
}
