package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"unicode"
)

func main() {

	// ignores := []string{
	// 	".git",
	// }
	// ext := []string{
	// 	".txt",
	// }
	// havingAWalk(".", ext, ignores)
	parsingTodos()
}

func havingAWalk(dir string, ext []string, ignores []string) {
	files := []string{}

	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {

		if err != nil {
			return err
		}

		if d.IsDir() {
			if slices.Contains(ignores, d.Name()) {
				return filepath.SkipDir
			}
		} else {
			e := filepath.Ext(path)
			if !(slices.Contains(ext, e)) {
				return nil
			}

			a, _ := filepath.Abs(path)
			files = append(files, a)
		}
		return nil
	})

	for _, f := range files {
		fmt.Println(f)
	}
}

func parsingTodos() {
	handle, err := os.OpenFile("test.txt", os.O_RDONLY, 0400)
	if err != nil {
		panic(err)
	}
	defer handle.Close()

	scan := bufio.NewScanner(handle)
	scan.Split(bufio.ScanLines)

	linenum := 1

	for scan.Scan() {
		line := scan.Text()
		if hasLineTodo(filepath.Ext(handle.Name()), line) {
			fmt.Printf("trouvé a la ligne %d \n", linenum)
		}

		if scan.Err() != nil {
			panic(scan.Err())
		}
		linenum++
	}

	// panic("Not implemented")
}

func hasLineTodo(ext string, s string) bool {
	comments := map[string]string{
		".sql": "--",
		".py":  "#",
		".nim": "#",
	}

	commentSymbol := "//"
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

func parse(s string) {
	panic("Not implemented")
}
