package main

import (
	"TODOS_Logger/internal/console"
	"TODOS_Logger/internal/exectime"
	"TODOS_Logger/internal/help"
	"TODOS_Logger/internal/logger"
	"TODOS_Logger/internal/todos"
	"TODOS_Logger/internal/walker"
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"
)

type set[T comparable] = map[T]struct{}

func main() {
	exectime.Benchmark(todoLogger)
}

func todoLogger() {

	args := os.Args[1:]
	if len(args) <= 0 {
		console.Printcln(console.RED, "Error: missing args !\n")
		help.Display()
		return
	}

	if args[0] == "-h" {
		help.Display()
		return
	}

	splet := strings.Split(args[0], ",")

	if len(splet) <= 0 {
		console.Printcln(console.RED, "Missing files to search for TODOS comments")
		return
	}

	exts := set[string]{}
	sb := strings.Builder{}
	for _, s := range splet {
		sb.WriteRune('.')
		sb.WriteString(s)
		exts[sb.String()] = struct{}{}
		sb.Reset()
	}

	fmt.Print("Looking for files : ")
	fmt.Println(strings.Join(slices.Collect(maps.Keys(exts)), ", "))

	ignores := set[string]{
		".git":         struct{}{},
		".vscode":      struct{}{},
		"bin":          struct{}{},
		"obj":          struct{}{},
		"node_modules": struct{}{},
		".venv":        struct{}{},
		"__pycache__":  struct{}{},
	}

	files, err := walker.WalkThisWay(".", exts, ignores)
	if err != nil {
		panic(err)
	}

	if len(files) <= 0 {
		console.Printcln(console.RED, "No file found")
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
		console.Printcln(console.RED, "No TODO comment found in files")
		return
	}

	log := logger.CreateLog(todoLines) //todo maybe be unused for md file or html

	if len(args) > 1 {
		options := args[1:]

		if slices.Contains(options, "c") {
			logger.WriteToConsole(log)
		}
		if slices.Contains(options, "f") {
			err := logger.WriteToFile(log)
			if err != nil {
				panic(err)
			}
		}
		if slices.Contains(options, "md") {
			log = logger.CreateLogToMd(todoLines)
			err := logger.WriteToSpecialFile(log, "md")
			if err != nil {
				panic(err)
			}
		}
		if slices.Contains(options, "html") {
			log := logger.CreateLogToHTML(todoLines)
			css := `
			<style>
				body{
					background-color: #3f3f3f;
				}
				div{
					background-color: #5e5e5e;
					border: 1px solid #7a7979; 
					padding: 10px; 
					margin:10px; 
					border-radius:5px;
				}
			</style>
			`
			sb := strings.Builder{}
			sb.WriteString(`
			<html>
			<head>`)
			sb.WriteString(css)
			sb.WriteString(`
			</head>
			<body>
			`)
			sb.WriteString(log)
			sb.WriteString(`
			</body>
			</html>
			`)

			logger.WriteToSpecialFile(sb.String(), "html")
		}

	} else {
		logger.WriteToConsole(log)
	}
}
