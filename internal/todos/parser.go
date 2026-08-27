package todos

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"unicode"
)

//opens file, find lines, caches them

// Checks if a line has a todo comment in it
func HasLineTodo(ext string, s string) bool {
	comments := map[string]string{
		".sql": "--",
		".py":  "#",
		".nim": "#",
	}

	commentSymbol := "//" //default val
	if _, ok := comments[ext]; ok {
		commentSymbol = comments[ext]
	}

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

// Opens a file, look for TODOS comments and theirs trailing comments.
// Returns all lines founds and closes the file.
func GetAll(file string) {
	handle, err := os.OpenFile(file, os.O_RDONLY, 0400)
	if err != nil {
		panic(err)
	}
	defer handle.Close()

	scan := bufio.NewScanner(handle)
	scan.Split(bufio.ScanLines)

	linenum := 1

	for scan.Scan() {
		line := scan.Text()
		if HasLineTodo(filepath.Ext(handle.Name()), line) {
			fmt.Printf("trouvé a la ligne %d \n", linenum)
		}

		if scan.Err() != nil {
			panic(scan.Err())
		}
		linenum++
	}

}
