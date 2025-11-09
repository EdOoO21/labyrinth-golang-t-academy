package application

import (
	"io"
	"math/rand"
	"testing"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

type fakeReader struct {
	gotMaze bool
}

func (f *fakeReader) GetMazeFromFile(path string) (*domain.Maze, error) {
	maze, _ := domain.NewMazeFromScheme([][]domain.CellType{
		{domain.CellWall, domain.CellWall, domain.CellWall, domain.CellWall, domain.CellWall},
		{domain.CellWall, domain.CellEmpty, domain.CellEmpty, domain.CellEmpty, domain.CellWall},
		{domain.CellWall, domain.CellEmpty, domain.CellEmpty, domain.CellEmpty, domain.CellWall},
		{domain.CellWall, domain.CellEmpty, domain.CellEmpty, domain.CellEmpty, domain.CellWall},
		{domain.CellWall, domain.CellWall, domain.CellWall, domain.CellWall, domain.CellWall},
	})

	f.gotMaze = true
	return maze, nil
}

type fakeWriter struct {
	printedMaze bool
	printedPath bool
}

func (f *fakeWriter) PrintMaze(path string, buf io.Writer, maze *domain.Maze) error {
	f.printedMaze = true
	return nil
}

func (f *fakeWriter) PrintPath(output string, path *domain.Path, maze *domain.Maze) error {
	f.printedPath = true
	return nil
}

func TestRun(t *testing.T) {

	tests := []struct {
		name            string
		input           *InputParams
		rand            *fakeRandom
		reader          *fakeReader
		writer          *fakeWriter
		wantErr         bool
		wantPrintedMaze bool
		wantPrintedPath bool
		wantGotMaze     bool
	}{
		{
			name:            "generate-dfs",
			input:           &InputParams{cmd: "generate", algorithm: "dfs", height: 5, width: 5, output: "out"},
			rand:            &fakeRandom{r: rand.New(rand.NewSource(42))},
			reader:          &fakeReader{false},
			writer:          &fakeWriter{false, false},
			wantErr:         false,
			wantPrintedMaze: true,
			wantPrintedPath: false,
			wantGotMaze:     false,
		},
		{
			name:            "generate-prim",
			input:           &InputParams{cmd: "generate", algorithm: "prim", height: 6, width: 6, output: "out"},
			rand:            &fakeRandom{r: rand.New(rand.NewSource(7))},
			reader:          &fakeReader{false},
			writer:          &fakeWriter{false, false},
			wantErr:         false,
			wantPrintedMaze: true,
			wantPrintedPath: false,
			wantGotMaze:     false,
		},
		{
			name:            "generate-bad-algo",
			input:           &InputParams{cmd: "generate", algorithm: "boom", height: 5, width: 5},
			rand:            &fakeRandom{r: rand.New(rand.NewSource(1))},
			reader:          &fakeReader{false},
			writer:          &fakeWriter{false, false},
			wantErr:         true,
			wantPrintedMaze: false,
			wantPrintedPath: false,
			wantGotMaze:     false,
		},
		{
			name:            "generate-small",
			input:           &InputParams{cmd: "generate", algorithm: "dfs", height: 1, width: 1},
			rand:            &fakeRandom{r: rand.New(rand.NewSource(123))},
			reader:          &fakeReader{false},
			writer:          &fakeWriter{false, false},
			wantErr:         false,
			wantPrintedMaze: true,
			wantPrintedPath: false,
			wantGotMaze:     false,
		},
		{
			name: "solve-astar-valid",
			input: &InputParams{
				cmd:        "solve",
				algorithm:  "astar",
				mazeFile:   "testdata/maze_small.txt",
				startPoint: domain.NewCell(1, 1),
				endPoint:   domain.NewCell(3, 3),
			},
			rand:            &fakeRandom{r: rand.New(rand.NewSource(42))},
			reader:          &fakeReader{false},
			writer:          &fakeWriter{false, false},
			wantErr:         false,
			wantPrintedPath: true,
			wantPrintedMaze: false,
			wantGotMaze:     true,
		},
		{
			name: "solve_invalid_algorithm",
			input: &InputParams{
				cmd:        "solve",
				algorithm:  "invalid",
				mazeFile:   "testdata/maze_small.txt",
				startPoint: domain.NewCell(1, 1),
				endPoint:   domain.NewCell(3, 3),
			},
			rand:            &fakeRandom{r: rand.New(rand.NewSource(42))},
			reader:          &fakeReader{false},
			writer:          &fakeWriter{false, false},
			wantErr:         true,
			wantPrintedPath: false,
			wantPrintedMaze: false,
			wantGotMaze:     false,
		},
		{
			name: "solve-dijkstra-valid",
			input: &InputParams{
				cmd:        "solve",
				algorithm:  "dijkstra",
				mazeFile:   "testdata/maze_small.txt",
				startPoint: domain.NewCell(1, 1),
				endPoint:   domain.NewCell(3, 3),
			},
			rand:            &fakeRandom{r: rand.New(rand.NewSource(99))},
			reader:          &fakeReader{false},
			writer:          &fakeWriter{false, false},
			wantErr:         false,
			wantPrintedPath: true,
			wantPrintedMaze: false,
			wantGotMaze:     true,
		},
		{
			name: "solve-start-oob",
			input: &InputParams{
				cmd:        "solve",
				algorithm:  "astar",
				mazeFile:   "testdata/maze_small.txt",
				startPoint: domain.NewCell(999, 999),
				endPoint:   domain.NewCell(3, 3),
			},
			rand:            &fakeRandom{r: rand.New(rand.NewSource(3))},
			reader:          &fakeReader{false},
			writer:          &fakeWriter{false, false},
			wantErr:         true,
			wantPrintedPath: false,
			wantPrintedMaze: false,
			wantGotMaze:     false,
		},
		{
			name: "solve-end-oob",
			input: &InputParams{
				cmd:        "solve",
				algorithm:  "astar",
				mazeFile:   "testdata/maze1.txt",
				startPoint: domain.NewCell(1, 1),
				endPoint:   domain.NewCell(999, 999),
			},
			rand:            &fakeRandom{r: rand.New(rand.NewSource(4))},
			reader:          &fakeReader{false},
			writer:          &fakeWriter{false, false},
			wantErr:         true,
			wantPrintedPath: false,
			wantPrintedMaze: false,
			wantGotMaze:     false,
		},
		{
			name:            "unknown-command",
			input:           &InputParams{cmd: "foobar"},
			rand:            &fakeRandom{r: rand.New(rand.NewSource(5))},
			reader:          &fakeReader{false},
			writer:          &fakeWriter{false, false},
			wantErr:         true,
			wantPrintedMaze: false,
			wantPrintedPath: false,
			wantGotMaze:     false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := Run(tc.input, tc.reader, tc.writer, tc.rand)
			if (err != nil) != tc.wantErr {
				t.Fatalf("Run() error = %v, wantErr = %v", err, tc.wantErr)
			}

			if tc.wantPrintedMaze && !tc.writer.printedMaze {
				t.Errorf("expected maze to be printed but it wasn't")
			}

			if tc.wantPrintedPath && !tc.writer.printedPath {
				t.Errorf("expected path to be printed but it wasn't")
			}

			if tc.wantGotMaze && !tc.reader.gotMaze {
				t.Errorf("expected solver to receive maze but it didn't")
			}
		})
	}
}
