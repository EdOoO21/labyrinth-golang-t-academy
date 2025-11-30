package parsers

import cell "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/cell"

type InputParams struct {
	Width, Height                    int
	Cmd, Algorithm, MazeFile, Output string
	StartPoint, EndPoint             cell.Cell
}
