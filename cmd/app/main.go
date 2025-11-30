package main

import (
	"fmt"
	"os"

	app "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/application/core"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/infrastructure/parsers"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/infrastructure/random"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/infrastructure/reader"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/infrastructure/writer"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("not enough arguments: enter generate or solve")
		os.Exit(1)
	}
	r := &reader.ConsoleReader{}
	w := &writer.ConsoleWriter{}
	rand := &random.RandomNumber{}
	input := &parsers.InputParams{}

	switch os.Args[1] {
	case "generate":
		exitOnError(parsers.GeneratorParseFlags(input, os.Args[2:]))

	case "solve":
		exitOnError(parsers.SolverParseFlags(input, os.Args[2:]))

	case "--help":
		writer.Helper()
		os.Exit(0)
	default:
		fmt.Println("incorrect first argument: should be generate or solve")
		os.Exit(1)
	}

	exitOnError(app.Run(input, r, w, rand))
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
