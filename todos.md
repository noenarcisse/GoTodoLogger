# todos

goroutine / folder ?
tests

err management a clean

# refacto & cleanup
walker.go : 
26 : if slices.Contains(ignores, d.Name()) { 
-> passer en set / map pour O(1) le contains

parser.go :
func IsCommentLine(ext string, line string) bool {
-> bugged

# maybe
maper par file ?