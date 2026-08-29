package todos

import "testing"

func Test_getCommentSymbol(t *testing.T) {
	ext := ".sql"

	res := getCommentSymbol(ext)
	exp := "--"

	if res != exp {
		t.Errorf(`getCommentSymbol(%s). 
		Result: %s
		Expected: %s`,
			ext, res, exp)
	}
}
func Test_getCommentSymbol2(t *testing.T) {
	ext := ".cs"

	res := getCommentSymbol(ext)
	exp := "//"

	if res != exp {
		t.Errorf(`getCommentSymbol(%s). 
		Result: %s
		Expected: %s`,
			ext, res, exp)
	}
}

func Test_getCommentSymbol3(t *testing.T) {
	ext := ".py"

	res := getCommentSymbol(ext)
	exp := "#"

	if res != exp {
		t.Errorf(`getCommentSymbol(%s). 
		Result: %s
		Expected: %s`,
			ext, res, exp)
	}
}

func Test_HasLineTodo(t *testing.T) {
	ext := ".py"
	code := `
	# un comment pas intéressant
	print("Hello")
	# todo pour tester
	print("World")

	#TODO
	print("Ca s'arrete jamais avec ces todos")
`

	res := HasLineTodo(ext, code)
	exp := true

	if res != exp {
		t.Errorf(`HasLineTodo(%s, %s). 
		Result: %t
		Expected: %t`,
			ext, code, res, exp)
	}
}

func Test_HasLineTodo2(t *testing.T) {
	ext := ".py"
	code := `
	# un comment pas intéressant
	print("Hello") # un autre
	print("Ca s'arrete jamais avec ces todos") #to ah non do
`

	res := HasLineTodo(ext, code)
	exp := false

	if res != exp {
		t.Errorf(`HasLineTodo(%s, %s). 
		Result: %t
		Expected: %t`,
			ext, code, res, exp)
	}
}

func Test_IsCommentLine(t *testing.T) {
	ext := ".py"
	line := "# un comment pas intéressant"

	res := IsCommentLine(ext, line)
	exp := true

	if res != exp {
		t.Errorf(`HasLineTodo(%s, %s). 
		Result: %t
		Expected: %t`,
			ext, line, res, exp)
	}
}

func Test_IsCommentLine2(t *testing.T) {
	ext := ".py"
	line := "         	# # un comment pas intéressant"

	res := IsCommentLine(ext, line)
	exp := true

	if res != exp {
		t.Errorf(`HasLineTodo(%s, %s). 
		Result: %t
		Expected: %t`,
			ext, line, res, exp)
	}
}

//todo
//bug found
//any matching comment-like char is considered
//its okay for the case code # comment
// but comment chars car appears in other places #hastag in py
// math for sql (10 - -1) =
func Test_IsCommentLine3(t *testing.T) {
	ext := ".py"
	line := "print(\"un code normal # avec un hashtag pour #twitter\")"

	res := IsCommentLine(ext, line)
	exp := false

	if res != exp {
		t.Errorf(`HasLineTodo(%s, %s). 
		Result: %t
		Expected: %t`,
			ext, line, res, exp)
	}
}
