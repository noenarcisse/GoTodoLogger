package help

import "fmt"

func Display() {
	fmt.Println("TODO Logger :")
	fmt.Println("todos [files] [outputs]")
	fmt.Println("\n	[files]")
	fmt.Println("		files can have several file extensions separated by a comma")
	fmt.Println("\n	[outputs]")
	fmt.Println("		outputs: where to print the logs")
	fmt.Println("		outputs: f -> file")
	fmt.Println("		outputs: md -> .md file")
	fmt.Println("		outputs: html -> .html file")
	fmt.Println("		outputs: c -> console")
	fmt.Println("		outputs can be accumulated")
	fmt.Println("		for example: todos txt md f c")
	fmt.Println("		by default, output write to the console")
}
