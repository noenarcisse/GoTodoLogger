package main

import (
	ae "TODOS_Logger/internal/app_error"
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
	"time"
)

type set[T comparable] = map[T]struct{}

func main() {
	exectime.Benchmark(todoLogger)
}

func todoLogger() {

	args := os.Args[1:]
	if len(args) <= 0 {
		console.Printcln(console.RED, ae.ErrMsg(ae.MISSING_ARGS))
		help.Display()
		return
	}

	if args[0] == "-h" {
		help.Display()
		return
	}

	splet := strings.Split(args[0], ",")

	if len(splet) <= 0 {
		console.Printcln(console.RED, ae.ErrMsg(ae.NO_FILE_SEARCH_SPECIFIED))
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
		console.Printcln(console.RED, ae.ErrMsg(ae.NO_FILE_FOUND))
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
		console.Printcln(console.RED, ae.ErrMsg(ae.NO_TODO_COMMENT))
		return
	}

	//todo refacto mieux ca en lazy ou equivalent ou prep un bloc d'ecriture separé de la console
	log := logger.CreateLog(todoLines) //todo maybe be unused for md file or html
	t := time.Now()
	logfilename := fmt.Sprintf("log_%d", t.Unix())

	if len(args) > 1 {
		options := args[1:] // todo go with a set

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
			err := logger.WriteToSpecialFile(log, logfilename, "md")
			if err != nil {
				panic(err)
			}
		}
		if slices.Contains(options, "html") {

			log := logger.CreateLogToHTML(todoLines)
			css := `<style>
/* 
prep inline CSS here
*/

body{
    background-color: #202224;
    font-family: Verdana, Geneva, Tahoma, sans-serif;
    font-size: 12px;
    color: white;
}
div{
    /* background-color: #3d3d3d; */
 
    padding: 10px; 
    margin:10px;

    border: 1px solid #7a7979;
    border-radius:5px;
}
a {
    font-family: Verdana, Geneva, Tahoma, sans-serif;
    color: #50b7e0;
    text-decoration: none;
}
a:hover{
    text-decoration: underline;
}
code {
    display: block;

    margin: 1em;
    padding-bottom:1em;
    padding-left: 1em;

    color:rgb(116, 164, 88);
    background-color: #333232;

    border-radius:5px;
}
</style>`
			sb := strings.Builder{}
			sb.WriteString(`
			<html>
			<head>`)
			sb.WriteString(css)
			sb.WriteString(`
			</head>
			<body>
			<div>
			<h1>`)
			sb.WriteString(logfilename)
			sb.WriteString(`</h1>
			</div>
			`)
			sb.WriteString(log)
			sb.WriteString("</body></html>")

			logger.WriteToSpecialFile(sb.String(), logfilename, "html")
		}

	} else {
		logger.WriteToConsole(log)
	}
}
