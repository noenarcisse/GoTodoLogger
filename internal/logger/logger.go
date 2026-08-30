package logger

import (
	"TODOS_Logger/internal/todos"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Create a full tetx log based on the TODOS found
func CreateLog(tds []todos.TodoLine) string {
	sb := strings.Builder{}

	for _, td := range tds {
		sb.WriteString(td.Format())
		sb.WriteString("\n")
	}

	return sb.String()
}

func WriteToConsole(s string) {
	fmt.Println(s)
}

func WriteToFile(s string) error {
	t := time.Now()

	logfilename := fmt.Sprintf("log_%d", t.Unix())
	fmt.Printf("Logging in %s\n", logfilename)
	dirname := "logs"

	err := os.MkdirAll(dirname, 0755)
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

	fmt.Printf("%d bytes ecrits\n", written)
	return nil
}
