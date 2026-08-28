package main

import (
	"GoLogTodos/internal/logger"
	"GoLogTodos/internal/todos"
	"GoLogTodos/internal/walker"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	bench(todoLogger)
}
func bench(f func()) {
	t1 := time.Now()
	f()
	execTime := time.Since(t1)
	fmt.Printf("exec time : %v\n", execTime)
}

func todoLogger() {

	args := os.Args[1:]

	if args[0] == "-h" {
		//display help
		//quits
	}

	splet := strings.Split(args[0], ",")

	if len(splet) <= 0 {
		panic("Aie")
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

	logger.WriteToFile(todoLines)

}
