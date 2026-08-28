package todos

import (
	"bufio"
	"os"
	"path/filepath"

	"strings"
	"unicode"
)

//opens file, find lines, caches them

func getCommentSymbol(ext string) string {
	comments := map[string]string{
		".sql": "--",
		".py":  "#",
		".nim": "#",
	}

	commentSymbol := "//" //default val
	if _, ok := comments[ext]; ok {
		commentSymbol = comments[ext]
	}

	return commentSymbol
}

// Checks if a line has a todo comment in it
func HasLineTodo(ext string, s string) bool {
	commentSymbol := getCommentSymbol(ext)
	todo := []rune(commentSymbol + "todo")
	index := 0
	for _, c := range s {
		if unicode.IsSpace(c) {
			continue
		}

		if unicode.ToLower(c) == todo[index] {
			index++

			if index == len(todo) {
				return true
			}
		} else {
			index = 0
		}
	}
	return false
}

func IsCommentLine(ext string, line string) bool {

	commentSymbol := getCommentSymbol(ext)
	commentRunes := []rune(commentSymbol)
	index := 0
	runes := []rune(line)

	for _, c := range runes {
		if unicode.IsSpace(c) {
			continue
		}
		if c == commentRunes[index] {
			index++
			if index == len(commentRunes) {
				return true
			}
		}
	}
	return false
}

// Opens a file, look for TODOS comments and theirs trailing comments.
// Returns all lines founds and closes the file.
func GetAll(file string) []TodoLine {
	todosFound := []TodoLine{}

	handle, err := os.OpenFile(file, os.O_RDONLY, 0400)
	if err != nil {
		panic(err)
	}
	defer handle.Close()

	scan := bufio.NewScanner(handle)
	scan.Split(bufio.ScanLines)

	linenum := 1
	todo := TodoLine{
		File:  file,
		Lines: map[int]string{},
	}
	sb := strings.Builder{}

	isTrailingLine := false
	ext := filepath.Ext(file)

	for scan.Scan() {

		line := scan.Text()

		if isTrailingLine {
			if IsCommentLine(ext, line) {

				todo.Lines[linenum] = strings.TrimSpace(line)
				// todo.TrailingLine = append(todo.TrailingLine, linenum)
				// sb.WriteString("\n")
				// sb.WriteString(strings.TrimSpace(line))

			} else {
				//si non ->
				todo.Content = sb.String()
				sb.Reset()
				todosFound = append(todosFound, todo)
				isTrailingLine = false
				todo = TodoLine{
					File:  file,
					Lines: map[int]string{},
				}
			}
		}

		if HasLineTodo(ext, line) {
			todo.Lines[linenum] = strings.TrimSpace(line)
			todo.LineNum = linenum
			// sb.WriteString(strings.TrimSpace(line))
			// sb.WriteString("\n")
			isTrailingLine = true
		}

		if scan.Err() != nil {
			panic(scan.Err())
		}
		linenum++
	}
	return todosFound
}
