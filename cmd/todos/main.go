package main

import (
	"TODOS_Logger/internal/exectime"
	"TODOS_Logger/internal/help"
	"TODOS_Logger/internal/logger"
	"TODOS_Logger/internal/todos"
	"TODOS_Logger/internal/walker"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	exectime.Benchmark(todoLogger)
}

func todoLogger() {
	const (
		Reset  = "\033[0m"
		Red    = "\033[31m"
		Green  = "\033[32m"
		Yellow = "\033[33m"
		Blue   = "\033[34m"
	)

	args := os.Args[1:]
	if len(args) <= 0 {
		fmt.Printf(Red + "Error: missing args !\n\n" + Reset)
		help.Display()
		return
	}

	if args[0] == "-h" {
		help.Display()
		return
	}

	splet := strings.Split(args[0], ",")

	if len(splet) <= 0 {
		fmt.Println(Red + "Missing files to search for TODOS comments" + Reset)
		return
	}

	exts := []string{}

	for _, s := range splet {
		exts = append(exts, "."+s) //todo fix ca
	}

	fmt.Print("Looking for files : ")
	fmt.Println(strings.Join(exts, ", "))

	ignores := []string{
		".git",
		".vscode",
		"bin",
		"obj",
		"node_modules",
		".venv", //uh?
		"__pycache__",
	}

	files, err := walker.WalkThisWay(".", exts, ignores)
	if err != nil {
		panic(err)
	}

	if len(files) <= 0 {
		fmt.Println(Red + "No files found" + Reset)
		return
	}

	todoLines := []todos.TodoLine{}

	for _, f := range files {
		fmt.Println(f)

		found := todos.GetAll(f)
		if len(found) > 0 {
			todoLines = append(todoLines, found...)
		}
	}

	if len(todoLines) <= 0 {
		fmt.Println(Red + "No TODOS found in files" + Reset)
		return
	}

	log := logger.CreateLog(todoLines)

	if len(args) > 1 {
		splet = strings.Split(args[1], "")

		if slices.Contains(splet, "c") {
			logger.WriteToConsole(log)
		}
		if slices.Contains(splet, "f") {
			err := logger.WriteToFile(log)
			if err != nil {
				panic(err)
			}
		}
	} else {
		logger.WriteToConsole(log)
	}
}
