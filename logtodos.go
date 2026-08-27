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
	for _, f := range files {
		fmt.Println(f)
		todos.GetAll(f)
	}

	// todos.GetAll("test.sql")
}
