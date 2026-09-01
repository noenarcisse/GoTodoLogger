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
