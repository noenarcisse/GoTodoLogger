package logger

import (
	"TODOS_Logger/internal/formatter"
	"TODOS_Logger/internal/todos"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// todo wip
type FormatterFunc func(todos.TodoLine) string

func createLog(tds []todos.TodoLine, f FormatterFunc) string {
	sb := strings.Builder{}
	for _, td := range tds {
		sb.WriteString(f(td))
		sb.WriteString("\n")
	}
	return sb.String()
}

// Create a full text log based on the TODOS found
func CreateLog(tds []todos.TodoLine) string {
	return createLog(tds, formatter.ToConsole)
}
func CreateLogToMd(tds []todos.TodoLine) string {
	return createLog(tds, formatter.ToMd)
}
func CreateLogToHTML(tds []todos.TodoLine) string {
	return createLog(tds, formatter.ToHTML)
}

func createLogDir(dirname string) error {

	err := os.MkdirAll(dirname, 0755)
	if err != nil {
		return err
	}
	return nil
}
func WriteToConsole(s string) {
	fmt.Println(s)
}
func WriteToFile(s string) error {
	t := time.Now()

	logfilename := fmt.Sprintf("log_%d", t.Unix())
	fmt.Printf("Logging in %s\n", logfilename)

	dirname := "logs"
	err := createLogDir(dirname)
	if err != nil {
		return err
	}

	handle, err := os.OpenFile(filepath.Join(dirname, logfilename), os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer handle.Close()

	written, err := handle.WriteString(s)
	if err != nil {
		return err
	}

	fmt.Printf("%d bytes written\n", written)
	return nil
}
func WriteToSpecialFile(s string, logfilename string, ext string) error {
	filetowrite := logfilename + "." + ext
	fmt.Printf("Logging in %s\n", filetowrite)

	dirname := "logs"
	err := createLogDir(dirname)
	if err != nil {
		return err
	}

	handle, err := os.OpenFile(filepath.Join(dirname, filetowrite), os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer handle.Close()

	written, err := handle.WriteString(s)
	if err != nil {
		return err
	}

	fmt.Printf("%d bytes written\n", written)
	return nil
}
