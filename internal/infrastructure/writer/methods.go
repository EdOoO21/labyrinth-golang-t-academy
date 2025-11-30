package writer

import (
	"fmt"
	"io"
	"os"

	cells "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/cell"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/path"
)

// Maze funcs ------------------------------------
func (с *ConsoleWriter) PrintMaze(path string, buf io.Writer, maze *maze.Maze) error {
	if path != "" {
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()
		buf = file
	}

	height, width := maze.GetHeight(), maze.GetWidth()
	for i := range height {
		for j := range width {
			cell, err := maze.GetCell(i, j)
			if err != nil {
				return err
			}
			if _, err := fmt.Fprintf(buf, "%s", string(cell)); err != nil {
				return err
			}
		}
		if _, err := fmt.Fprintf(buf, "\n"); err != nil {
			return err
		}
	}
	return nil
}

// Path funcs ------------------------------------

func (c *ConsoleWriter) PrintPath(output string, path *path.Path, maze *maze.Maze) error {
	p := path.GetPath()

	for i, cell := range p {
		switch i {
		case 0:
			_ = maze.SetCell(cell.Row, cell.Col, cells.CellStart)
		case len(p) - 1:
			_ = maze.SetCell(cell.Row, cell.Col, cells.CellEnd)
		default:
			_ = maze.SetCell(cell.Row, cell.Col, cells.CellPath)
		}
	}
	err := c.PrintMaze(output, os.Stdout, maze)

	return err
}

// Other ---------------------------------------------------

func Helper() {
	fmt.Println("Usage: maze-app [-hV] [COMMAND]\nMaze generator and solver CLI application.\n  -h, --help      Show this help message and exit.\n  -V, --version   Print version information and exit.\nCommands:\n  generate  Generate a maze with specified algorithm and dimensions.\n  solve     Solve a maze with specified algorithm and points.")
}
