package support

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/cell"
)

type PqItem struct {
	Dist int
	C    cell.Cell
}

type AstarItem struct {
	Dist int
	H    int
	C    cell.Cell
}

type PQDijkstra []PqItem

type PQAStar []AstarItem
