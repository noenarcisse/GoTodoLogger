package logger

import (
	"TODOS_Logger/internal/formatter"
	"TODOS_Logger/internal/todos"
	"html"
	"os"
	"strings"
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

func PrepareHTMLContent(files []string, templateHtml string, css string, log string, logfilename string) string {
	sb := strings.Builder{}
	for _, f := range files {
		sb.WriteString("<li>")
		sb.WriteString(html.EscapeString(f))
		sb.WriteString("</li>")
	}
	html2 := strings.Replace(templateHtml, "[css]", css, 1)
	html2 = strings.Replace(html2, "[log]", log, 1)
	html2 = strings.ReplaceAll(html2, "[filename]", strings.ToUpper(logfilename))
	html2 = strings.Replace(html2, "[files]", sb.String(), 1)
	html2 = strings.Replace(html2, "[log]", log, 1)
	return html2
}
