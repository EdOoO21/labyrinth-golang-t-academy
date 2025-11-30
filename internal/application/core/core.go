package core

import (
	"fmt"
	"os"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/application/generators/dfs"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/application/generators/prim"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/application/models"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/application/solvers/astar"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/application/solvers/dijkstra"
	interfaces "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/interfaces"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/infrastructure/parsers"
)

func Run(input *parsers.InputParams, r models.Reader, w models.Writer, rand models.Random) error {
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
func generateMaze(input *parsers.InputParams, w models.Writer, rand models.Random) error {
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

func makeGenerator(input *parsers.InputParams, rand models.Random) interfaces.Generator {
	switch input.Algorithm {
	case "dfs":
		return &dfs.DFS{Rnd: rand}
	case "prim":
		return &prim.Prim{Rnd: rand}
	default:
		return nil
	}
}

// Solver --------------------------------------------------------------------------------------------------
func solveMaze(input *parsers.InputParams, w models.Writer, r models.Reader) error {
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

func makeSolver(input *parsers.InputParams) interfaces.Solver {
	switch input.Algorithm {
	case "astar":
		return &astar.AStar{}
	case "dijkstra":
		return &dijkstra.Dijkstra{}
	default:
		return nil
	}
}
