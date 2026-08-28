package todos

import (
	"fmt"
	"strconv"
	"strings"
)

type TodoLine struct {
	File         string //abs filepath ref
	Lines        map[int]string
	LineNum      int    // number of first line with todo
	TrailingLine []int  // numbers of the comments directly after the first todo line
	Content      string // lines content
}

func (tl TodoLine) String() string {
	return fmt.Sprintf(`FILE : 
Filepath : %s
Lines : %v
Lines : %s 
Line num : %d 
Trailing lines : %v`,
		tl.File, tl.Lines, tl.Content, tl.LineNum, tl.TrailingLine)
}

func (tl TodoLine) Format() string {
	sb := strings.Builder{}

	sb.WriteString(tl.File)
	sb.WriteString(" : \n")

	for k, v := range tl.Lines {
		sb.WriteString(strconv.Itoa(k)) // err ?
		sb.WriteString(" : ")
		sb.WriteString(v)
		sb.WriteString("\n")
	}
	return sb.String()
	// panic("Not implemented yet")
}
