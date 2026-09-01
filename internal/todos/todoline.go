package todos

import (
	sm "TODOS_Logger/pkg/sorted_map"
	"fmt"
)

type TodoLine struct {
	File    string                     //abs filepath ref
	Lines   *sm.SortedMap[int, string] // lazy sorted map
	LineNum int                        // number of first line with todo
}

func (tl TodoLine) String() string {
	return fmt.Sprintf(`FILE : 
Filepath : %s
Lines : %v
Todo Line num : %d`,
		tl.File, tl.Lines, tl.LineNum)
}

//todo remove this
// -> moved to formatter.go

// func (tl TodoLine) Format() string {
// 	sb := strings.Builder{}

// 	sb.WriteString(tl.File)
// 	sb.WriteString(":")
// 	sb.WriteString(strconv.Itoa(tl.LineNum))
// 	sb.WriteString(" : \n")

// 	for k, v := range tl.Lines.Items() {
// 		sb.WriteString(strconv.Itoa(k))
// 		sb.WriteString(" : ")
// 		sb.WriteString(v)
// 		sb.WriteString("\n")
// 	}
// 	return sb.String()
// }

// func (tl TodoLine) FormatToFile() string {
// 	sb := strings.Builder{}

// 	sb.WriteString(tl.File)
// 	sb.WriteString(":")
// 	sb.WriteString(strconv.Itoa(tl.LineNum))
// 	sb.WriteRune('\n')

// 	for k, v := range tl.Lines.Items() {
// 		sb.WriteString(strconv.Itoa(k))
// 		sb.WriteString(" : ")
// 		sb.WriteString(v)
// 		sb.WriteString("\n")
// 	}
// 	return sb.String()
// }

// func (tl TodoLine) FormatToMd() string {
// 	sb := strings.Builder{}

// 	filename := fmt.Sprintf("[%s:%s](vscode://file/%s:%s)", tl.File, strconv.Itoa(tl.LineNum), tl.File, strconv.Itoa(tl.LineNum))

// 	sb.WriteString(filepath.ToSlash(filename))
// 	sb.WriteString(" <br> \n")
// 	sb.WriteString("```\n")

// 	for k, v := range tl.Lines.Items() {
// 		sb.WriteString(strconv.Itoa(k))
// 		sb.WriteString(" : ")
// 		sb.WriteString(v)
// 		sb.WriteString("\n")
// 	}
// 	sb.WriteString("```")
// 	sb.WriteString("\n<br>\n")
// 	return sb.String()
// }

// func (tl TodoLine) FormatToHTML() string {
// 	sb := strings.Builder{}

// 	//vscode://
// 	filename := fmt.Sprintf("<a href=\"vscode://file/%s:%s\">%s:%s</a>", tl.File, strconv.Itoa(tl.LineNum), tl.File, strconv.Itoa(tl.LineNum))
// 	sb.WriteString("<div>\n")
// 	sb.WriteString("📄 ")
// 	sb.WriteString(filename)
// 	sb.WriteString("<br/>\n")
// 	sb.WriteString("<code>")

// 	for k, v := range tl.Lines.Items() {

// 		sb.WriteString("<br/>\n")
// 		sb.WriteString(strconv.Itoa(k))
// 		sb.WriteString("   ")
// 		sb.WriteString(v)
// 	}
// 	sb.WriteString("</code>")
// 	sb.WriteString("</div>\n")
// 	return sb.String()
// }
