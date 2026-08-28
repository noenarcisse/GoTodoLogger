package help

import "fmt"

func Display() {
	fmt.Println("TODO Logger :")
	fmt.Println("todos [files] [outputs]")
	fmt.Println("	file: can have several format separated by a comma")
	fmt.Println("	outputs: where to print the logs")
	fmt.Println("	outputs: f -> file")
	fmt.Println("	outputs: c -> console")
	fmt.Println("	outputs can be either f or c or fc or cf")
	fmt.Println("	by default, output write to the console")
}
