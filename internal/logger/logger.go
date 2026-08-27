package logger

import (
	"GoLogTodos/internal/todos"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

//log stuff in console or in a file ?
//requires time, guid etc ?

func WriteToFile(tds []todos.TodoLine) {
	t := time.Now()

	logfilename := fmt.Sprintf("log_%d", t.Unix())
	fmt.Println(logfilename)

	// s := "Un truc important"

	dirname := "logs"
	err := os.MkdirAll(dirname, 0755)
	if err != nil {
		panic(err)
	}

	handle, err := os.OpenFile(filepath.Join(dirname, logfilename), os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err) // rlly ?
	}
	defer handle.Close()

	sb := strings.Builder{}

	for _, td := range tds {
		sb.WriteString(td.String())
		sb.WriteString("\n\n")
	}

	written, err := handle.WriteString(sb.String()) //todo change this
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d bytes ecrits\n", written)

}

func WriteFileTest() {
	t := time.Now()

	logfilename := fmt.Sprintf("log_%d", t.Unix())
	fmt.Println(logfilename)

	s := "Un truc important"

	dirname := "logs"
	err := os.Mkdir(dirname, 0644)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filepath.Join(dirname, logfilename), []byte(s), 0644)
	if err != nil {
		panic(err)
	}
}
