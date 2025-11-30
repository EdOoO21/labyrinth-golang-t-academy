package maze

import (
	"strings"
	"testing"

	cell "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/cell"
)

func TestNewMazeFromReader(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		wantErr      bool
		wantH, wantW int
	}{
		{
			name:    "valid 3x3",
			input:   "###\n#.#\n###\n",
			wantErr: false,
			wantH:   3,
			wantW:   3,
		},
		{
			name:    "unequal rows",
			input:   "###\n##\n###\n",
			wantErr: true,
		},
		{
			name:    "empty input",
			input:   "",
			wantErr: true,
		},
		{
			name:    "only empty lines",
			input:   "\n\n\n",
			wantErr: true,
		},
		{
			name:    "single line",
			input:   "####\n",
			wantErr: false,
			wantH:   1,
			wantW:   4,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			m, err := NewMazeFromReader(strings.NewReader(tc.input))
			if (err != nil) != tc.wantErr {
				t.Fatalf("expected error=%v, got %v", tc.wantErr, err)
			}
			if !tc.wantErr {
				if m.height != tc.wantH || m.width != tc.wantW {
					t.Errorf("wrong size: got %dx%d, want %dx%d", m.height, m.width, tc.wantH, tc.wantW)
				}
			}
		})
	}
}

func TestNewMaze(t *testing.T) { // тут можно не писать больше тестов, потому что масштабирование одинаковое
	m := NewMaze(2, 3)
	if m.height != 5 || m.width != 7 {
		t.Errorf("unexpected dimensions: got %dx%d, want 5x7", m.height, m.width)
	}
	for i := 0; i < m.height; i++ {
		for j := 0; j < m.width; j++ {
			if m.scheme[i][j] != cell.CellWall {
				t.Errorf("expected CellWall at (%d,%d)", i, j)
			}
		}
	}
}

func TestNewMazeFromScheme(t *testing.T) { // тут тоже, эта функция используется только для теста
	// передаются в нее только корректные значения
	scheme := [][]cell.CellType{
		{cell.CellWall, cell.CellWall},
		{cell.CellWall, cell.CellEmpty},
	}
	m, err := NewMazeFromScheme(scheme)
	if err != nil {
		t.Fatal(err)
	}
	if m.height != 2 || m.width != 2 {
		t.Errorf("wrong size")
	}
}
