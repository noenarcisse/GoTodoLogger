package todos

import (
	sm "TODOS_Logger/pkg/sorted_map"
	"fmt"
	"strconv"
	"strings"
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

func (tl TodoLine) Format() string {
	sb := strings.Builder{}

	sb.WriteString(tl.File)
	sb.WriteString(" : \n")

	for k, v := range tl.Lines.Items() {
		sb.WriteString(strconv.Itoa(k))
		sb.WriteString(" : ")
		sb.WriteString(v)
		sb.WriteString("\n")
	}
	return sb.String()
}
