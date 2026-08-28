package todos

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

type TodoLine struct {
	File         string //abs filepath ref
	LineNum      int    // number of first line with todo
	TrailingLine []int  // numbers of the comments directly after the first todo line
	Content      string // lines content
}

func (tl TodoLine) String() string {
	return fmt.Sprintf(`FILE : 
Filepath : %s
Lines : %s 
Line num : %d 
Trailing lines : %v`,
		tl.File, tl.Content, tl.LineNum, tl.TrailingLine)
}

func (tl TodoLine) Format() string {
	panic("Not implemented yet")
}

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
	todo := TodoLine{File: file}
	sb := strings.Builder{}
	isTrailingLine := false
	ext := filepath.Ext(file)

	for scan.Scan() {

		line := scan.Text()

		if isTrailingLine {
			if IsCommentLine(ext, line) {
				todo.TrailingLine = append(todo.TrailingLine, linenum)
				sb.WriteString("\n")
				sb.WriteString(strings.TrimSpace(line))

			} else {
				//si non ->
				todo.Content = sb.String()
				sb.Reset()
				todosFound = append(todosFound, todo)
				isTrailingLine = false
				todo = TodoLine{File: file}
			}
		}

		if HasLineTodo(ext, line) {
			todo.LineNum = linenum
			sb.WriteString(strings.TrimSpace(line))
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
