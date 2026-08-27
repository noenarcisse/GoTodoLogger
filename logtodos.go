package main

import (
	"GoLogTodos/internal/todos"
	"GoLogTodos/internal/walker"
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]

	for _, a := range args {
		fmt.Println(a)
	}

	ignores := []string{
		".git",
		".vscode",
		"bin",
		"obj",
		"node_modules",
		".venv", //uh?
		"__pycache__",
	}
	ext := []string{
		".go",
	}
	files, err := walker.WalkThisWay(".", ext, ignores)
	if err != nil {
		panic(err)
	}

	todoLines := []todos.TodoLine{}

	for _, f := range files {
		fmt.Println(f)

		found := todos.GetAll(f)
		if len(found) > 0 {
			todoLines = append(todoLines, found...)
		}
	}

	for _, l := range todoLines {
		fmt.Println(l)
	}

}
