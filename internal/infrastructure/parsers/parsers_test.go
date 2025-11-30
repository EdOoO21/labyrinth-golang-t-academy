package parsers

import (
	"fmt"
	"testing"

	cell "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/cell"
)

func TestGeneratorParseFlags(t *testing.T) {
	tests := []struct {
		name     string
		input    *InputParams
		args     []string
		expected *InputParams
		err      error
	}{
		{
			name:  "missing width",
			input: &InputParams{},
			args:  []string{"-h", "5", "-a", "dfs"},
			expected: &InputParams{
				Cmd:       "generate",
				Width:     0,
				Height:    5,
				Algorithm: "dfs",
				Output:    "",
			},
			err: fmt.Errorf("--width=%d is not entered or is invalid: enter integer >0", 0),
		},
		{
			name:  "missing height",
			input: &InputParams{},
			args:  []string{"-w", "5", "-a", "prim"},
			expected: &InputParams{
				Cmd:       "generate",
				Width:     5,
				Height:    0,
				Algorithm: "prim",
				Output:    "",
			},
			err: fmt.Errorf("--height=%d is not entered or is invalid: enter integer >0", 0),
		},
		{
			name:  "negative width",
			input: &InputParams{},
			args:  []string{"-a", "dfs", "-w", "-22", "-h", "23"},
			expected: &InputParams{
				Cmd:       "generate",
				Width:     -22,
				Height:    23,
				Algorithm: "dfs",
				Output:    "",
			},
			err: fmt.Errorf("--width=%d is not entered or is invalid: enter integer >0", -22),
		},
		{
			name:  "negative height",
			input: &InputParams{},
			args:  []string{"-a", "prim", "-w", "12", "-h", "-2"},
			expected: &InputParams{
				Cmd:       "generate",
				Width:     12,
				Height:    -2,
				Algorithm: "prim",
				Output:    "",
			},
			err: fmt.Errorf("--height=%d is not entered or is invalid: enter integer >0", -2),
		},
		{
			name:  "missing algorithm",
			input: &InputParams{},
			args:  []string{"-w", "5", "-h", "5"},
			expected: &InputParams{
				Cmd:       "generate",
				Width:     5,
				Height:    5,
				Algorithm: "",
				Output:    "",
			},
			err: fmt.Errorf("--algorithm is empty: enter valid name"),
		},
		{
			name:  "valid long flags",
			input: &InputParams{},
			args:  []string{"--width", "3", "--height", "3", "--algorithm", "dfs", "--output", "out.txt"},
			expected: &InputParams{
				Cmd:       "generate",
				Width:     3,
				Height:    3,
				Algorithm: "dfs",
				Output:    "out.txt",
			},
			err: nil,
		},
		{
			name:  "valid short flags",
			input: &InputParams{},
			args:  []string{"-w", "2", "-h", "2", "-a", "prim", "-o", "f"},
			expected: &InputParams{
				Cmd:       "generate",
				Width:     2,
				Height:    2,
				Algorithm: "prim",
				Output:    "f",
			},
			err: nil,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := GeneratorParseFlags(tc.input, tc.args)

			if tc.err == nil && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if tc.err != nil {
				if err == nil {
					t.Fatalf("expected error %q but got nil", tc.err.Error())
				}
				if err.Error() != tc.err.Error() {
					t.Fatalf("expected error %q, got %q", tc.err.Error(), err.Error())
				}
			}

			if tc.expected != nil && tc.input != nil {
				got := tc.input
				want := tc.expected

				if got.Cmd != want.Cmd ||
					got.Width != want.Width ||
					got.Height != want.Height ||
					got.Algorithm != want.Algorithm ||
					got.Output != want.Output {
					t.Fatalf("unexpected parsed input\ngot:  %+v\nwant: %+v", got, want)
				}
			} else if (tc.expected == nil) != (tc.input == nil) {
				t.Fatalf("unexpected parsed input\ngot:  %+v\nwant: %+v", tc.input, tc.expected)
			}
		})
	}

}

func TestSolverParseFlags(t *testing.T) {
	tests := []struct {
		name     string
		input    *InputParams
		args     []string
		expected *InputParams
		err      error
	}{
		{
			name:  "missing algorithm",
			input: &InputParams{},
			args:  []string{"-f", "maze.txt", "-s", "0,0", "-e", "1,1"},
			expected: &InputParams{
				Cmd:        "solve",
				Algorithm:  "",
				MazeFile:   "maze.txt",
				Output:     "",
				StartPoint: cell.NewCell(0, 0),
				EndPoint:   cell.NewCell(0, 0),
			},
			err: fmt.Errorf("--algorithm is empty: enter valid name"),
		},
		{
			name:  "missing file",
			input: &InputParams{},
			args:  []string{"-a", "dijkstra", "-s", "0,0", "-e", "1,1"},
			expected: &InputParams{
				Cmd:        "solve",
				Algorithm:  "dijkstra",
				MazeFile:   "",
				Output:     "",
				StartPoint: cell.NewCell(0, 0),
				EndPoint:   cell.NewCell(0, 0),
			},
			err: fmt.Errorf("--file is empty: enter valid name"),
		},
		{
			name:  "invalid start format",
			input: &InputParams{},
			args:  []string{"-a", "dijkstra", "-f", "m", "-s", "-12,321", "-e", "1,1"},
			expected: &InputParams{
				Cmd:        "solve",
				Algorithm:  "dijkstra",
				MazeFile:   "m",
				Output:     "",
				StartPoint: cell.NewCell(0, 0),
				EndPoint:   cell.NewCell(0, 0),
			},
			err: fmt.Errorf("invalid point format: %s is invalid value", "(-12,321)"),
		},
		{
			name:  "invalid end format",
			input: &InputParams{},
			args:  []string{"-a", "dijkstra", "-f", "m", "-s", "11,321", "-e", "-1,21"},
			expected: &InputParams{
				Cmd:        "solve",
				Algorithm:  "dijkstra",
				MazeFile:   "m",
				Output:     "",
				StartPoint: cell.NewCell(11, 321),
				EndPoint:   cell.NewCell(0, 0),
			},
			err: fmt.Errorf("invalid point format: %s is invalid value", "(-1,21)"),
		},
		{
			name:  "negative coords",
			input: &InputParams{},
			args:  []string{"-a", "dijkstra", "-f", "m", "-s", "-1,0", "-e", "1,1"},
			expected: &InputParams{
				Cmd:        "solve",
				Algorithm:  "dijkstra",
				MazeFile:   "m",
				Output:     "",
				StartPoint: cell.NewCell(0, 0),
				EndPoint:   cell.NewCell(0, 0),
			},
			err: fmt.Errorf("invalid point format: (%d,%d) is invalid value", -1, 0),
		},
		{
			name:  "valid flags no output",
			input: &InputParams{},
			args:  []string{"-a", "astar", "-f", "m", "-s", "0,0", "-e", "1,1"},
			expected: &InputParams{
				Cmd:        "solve",
				Algorithm:  "astar",
				MazeFile:   "m",
				Output:     "",
				StartPoint: cell.NewCell(1, 1),
				EndPoint:   cell.NewCell(3, 3),
			},
			err: nil,
		},
		{
			name:  "valid flags with output",
			input: &InputParams{},
			args:  []string{"-o", "file", "-a", "dijkstra", "-f", "m", "-s", "2,44", "-e", "33,2113"},
			expected: &InputParams{
				Cmd:        "solve",
				Algorithm:  "dijkstra",
				MazeFile:   "m",
				Output:     "file",
				StartPoint: cell.NewCell(2, 44),
				EndPoint:   cell.NewCell(33, 2113),
			},
			err: nil,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := SolverParseFlags(tc.input, tc.args)

			if tc.err == nil && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if tc.err != nil {
				if err == nil {
					t.Fatalf("expected error %q but got nil", tc.err.Error())
				}
				if err.Error() != tc.err.Error() {
					t.Fatalf("expected error %q, got %q", tc.err.Error(), err.Error())
				}
			}

			if tc.expected != nil && tc.input != nil {
				got := tc.input
				want := tc.expected

				if got.Cmd != want.Cmd ||
					got.Width != want.Width ||
					got.Height != want.Height ||
					got.Algorithm != want.Algorithm ||
					got.Output != want.Output {
					t.Fatalf("unexpected parsed input\ngot:  %+v\nwant: %+v", got, want)
				}
			} else if (tc.expected == nil) != (tc.input == nil) {
				t.Fatalf("unexpected parsed input\ngot:  %+v\nwant: %+v", tc.input, tc.expected)
			}
		})
	}
}

func TestParseInputPoint(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		x, y int
		err  error
	}{
		{
			name: "valid input 0,0",
			arg:  "0,0",
			x:    1, y: 1,
			err: nil,
		},
		{
			name: "valid input 1,2",
			arg:  "1,2",
			x:    3, y: 5,
			err: nil,
		},
		{
			name: "missing comma",
			arg:  "12",
			x:    -1, y: -1,
			err: fmt.Errorf("12, expected format: x,y"),
		},
		{
			name: "too many parts",
			arg:  "1,2,3",
			x:    -1, y: -1,
			err: fmt.Errorf("1,2,3, expected format: x,y"),
		},
		{
			name: "non-integer x",
			arg:  "a,2",
			x:    -1, y: -1,
			err: fmt.Errorf("strconv.Atoi: parsing \"a\": invalid syntax"),
		},
		{
			name: "non-integer y",
			arg:  "1,b",
			x:    -1, y: -1,
			err: fmt.Errorf("strconv.Atoi: parsing \"b\": invalid syntax"),
		},
		{
			name: "negative x",
			arg:  "-1,2",
			x:    -1, y: -1,
			err: fmt.Errorf("(-1,2) is invalid value"),
		},
		{
			name: "negative y",
			arg:  "1,-3",
			x:    -1, y: -1,
			err: fmt.Errorf("(1,-3) is invalid value"),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			x, y, err := parseInputPoint(tc.arg)

			if err != nil && tc.err == nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if tc.err != nil {
				if err == nil {
					t.Fatalf("expected error: %v, got %v", tc.err, err)
				}
				if err.Error() != tc.err.Error() {
					t.Fatalf("expected error: %v, got %v", tc.err, err)
				}
			}

			if x != tc.x && y != tc.y {
				t.Fatalf("expected values (%v, %v), got (%v, %v)", tc.x, tc.y, x, y)
			}
		})
	}
}
