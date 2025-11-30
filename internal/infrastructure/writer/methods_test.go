package writer

import (
	"bytes"
	"testing"

	cell "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/cell"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
)

func TestPrintMaze(t *testing.T) {
	maze1, _ := maze.NewMazeFromScheme([][]cell.CellType{
		{cell.CellWall, cell.CellWall, cell.CellWall},
		{cell.CellWall, cell.CellWall, cell.CellWall},
		{cell.CellWall, cell.CellWall, cell.CellWall},
	})
	maze2, _ := maze.NewMazeFromScheme([][]cell.CellType{
		{cell.CellWall, cell.CellWall, cell.CellWall},
		{cell.CellWall, cell.CellEmpty, cell.CellWall},
		{cell.CellWall, cell.CellWall, cell.CellWall},
	})
	maze3, _ := maze.NewMazeFromScheme([][]cell.CellType{
		{cell.CellWall, cell.CellWall, cell.CellWall, cell.CellWall, cell.CellWall},
		{cell.CellWall, cell.CellEmpty, cell.CellEmpty, cell.CellEmpty, cell.CellWall},
		{cell.CellWall, cell.CellEmpty, cell.CellWall, cell.CellEmpty, cell.CellWall},
		{cell.CellWall, cell.CellEmpty, cell.CellEmpty, cell.CellEmpty, cell.CellWall},
		{cell.CellWall, cell.CellWall, cell.CellWall, cell.CellWall, cell.CellWall},
	})

	tests := []struct {
		name string
		maze *maze.Maze
		want string
	}{
		{
			name: "3x3 all walls",
			maze: maze1,
			want: "###\n###\n###\n",
		},
		{
			name: "3x3 center empty",
			maze: maze2,
			want: "###\n# #\n###\n",
		},
		{
			name: "2x2 scaled 5x5",
			maze: maze3,
			want: "#####\n#   #\n# # #\n#   #\n#####\n",
		},
	}

	writer := &ConsoleWriter{}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := writer.PrintMaze("", &buf, tc.maze)
			if err != nil {
				t.Fatalf("PrintMaze returned error: %v", err)
			}
			got := buf.String()
			if got != tc.want {
				t.Errorf("unexpected output:\nwant:\n%s\ngot:\n%s", tc.want, got)
			}
		})
	}
}
