package parsers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
	cell "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/cell"
)

func GeneratorParseFlags(input *InputParams, args []string) error {
	genFS := pflag.NewFlagSet("generate", pflag.ExitOnError)

	width := genFS.IntP("width", "w", 0, "maze width")
	height := genFS.IntP("height", "h", 0, "maze height")
	algo := genFS.StringP("algorithm", "a", "", "generate algorithm")
	output := genFS.StringP("output", "o", "", "output place")

	if err := genFS.Parse(args); err != nil {
		return err
	}

	input.Cmd = "generate"
	input.Width = *width
	input.Height = *height
	input.Algorithm = *algo
	input.Output = *output

	if input.Width <= 0 {
		return fmt.Errorf("--width=%d is not entered or is invalid: enter integer >0", input.Width)
	}

	if input.Height <= 0 {
		return fmt.Errorf("--height=%d is not entered or is invalid: enter integer >0", input.Height)
	}

	if input.Algorithm == "" {
		return fmt.Errorf("--algorithm is empty: enter valid name")
	}

	return nil
}

func SolverParseFlags(input *InputParams, args []string) error {
	solveFS := pflag.NewFlagSet("solve", pflag.ContinueOnError)

	algo := solveFS.StringP("algorithm", "a", "", "solve algorithm")
	mazeFile := solveFS.StringP("file", "f", "", "maze file")
	output := solveFS.StringP("output", "o", "", "output place")
	startStr := solveFS.StringP("start", "s", "", "start point")
	endStr := solveFS.StringP("end", "e", "", "end point")

	if err := solveFS.Parse(args); err != nil {
		return err
	}

	input.Cmd = "solve"
	input.Algorithm = *algo
	input.MazeFile = *mazeFile
	input.Output = *output

	if input.Algorithm == "" {
		return fmt.Errorf("--algorithm is empty: enter valid name")
	}

	if input.MazeFile == "" {
		return fmt.Errorf("--file is empty: enter valid name")
	}

	x, y, err := parseInputPoint(*startStr)
	if err != nil {
		return fmt.Errorf("invalid point format: %w", err)
	}
	input.StartPoint = cell.NewCell(x, y)

	x, y, err = parseInputPoint(*endStr)
	if err != nil {
		return fmt.Errorf("invalid point format: %w", err)
	}
	input.EndPoint = cell.NewCell(x, y)

	return nil
}

func parseInputPoint(s string) (int, int, error) {
	str := strings.Split(s, ",")
	if len(str) != 2 {
		return -1, -1, fmt.Errorf("%s, expected format: x,y", s)
	}
	var x, y int
	var err error

	if x, err = strconv.Atoi(str[0]); err != nil {
		return -1, -1, err
	}

	if y, err = strconv.Atoi(str[1]); err != nil {
		return -1, -1, err
	}

	if x < 0 || y < 0 {
		return -1, -1, fmt.Errorf("(%d,%d) is invalid value", x, y)
	}

	return 2*x + 1, 2*y + 1, nil
}
