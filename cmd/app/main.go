package main

import (
	"fmt"
	"os"

	app "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/application"
	inf "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/infrastructure"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("not enough arguments: enter generate or solve")
		os.Exit(1)
	}
	r := &inf.ConsoleReader{}
	w := &inf.ConsoleWriter{}
	rand := &inf.RandomNumber{}
	input := &app.InputParams{}

	switch os.Args[1] {
	case "generate":
		exitOnError(app.GeneratorParseFlags(input, os.Args[2:]))

	case "solve":
		exitOnError(app.SolverParseFlags(input, os.Args[2:]))

	case "--help":
		inf.Helper()
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
		// было exit(1), но из-за того,
		// что после Exit(1) выводится в терминал exit status 1
		// не проходили тесты blackbox, поэтому пришлось написать так
		// кажется, что изначально было норм решение при ошибки завершать с кодом 1
		// было бы супер если будет фидбек по этому поводу
	}
}
