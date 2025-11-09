package infrastructure

import (
	"os"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

func (с *ConsoleReader) GetMazeFromFile(path string) (*domain.Maze, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	maze, err := domain.NewMazeFromReader(file)
	if err != nil {
		return nil, err
	}
	return maze, nil
}

// тест нет смысла писать, потому что NewMazeFromReader и так покрыт тестом, а логики больше никакой тут нет
