package main

import (
	"GoLogTodos/internal/todos"
	"GoLogTodos/internal/walker"
)

func main() {

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
		".py",
	}
	walker.WalkThisWay(".", ext, ignores)
	todos.GetAll("test.sql")
}
