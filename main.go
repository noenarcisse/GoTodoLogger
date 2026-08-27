package main

import (
	"GoLogTodos/internal/todos"
	"GoLogTodos/internal/walker"
	"fmt"
	"os"
	"strings"
)

func main() {

	args := os.Args[1:]
	exts := []string{
		".go",
	}
	for _, a := range args {
		exts = append(exts, "."+a)
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
