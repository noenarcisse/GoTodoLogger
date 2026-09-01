package apperror

type Error int

const (
	UNKNOWN Error = iota
	MISSING_ARGS
	NO_FILE_SEARCH_SPECIFIED //user input issue
	NO_FILE_FOUND            // no file found with specifiie dinput rename this ?
	NO_TODO_COMMENT
)

func ErrMsg(e Error) string {
	return ErrMsg2(e, "")
}

func ErrMsg2(e Error, infos string) string {

	var msg string

	switch e {
	case MISSING_ARGS:
		msg = "Error: missing args !" + infos
	case NO_FILE_SEARCH_SPECIFIED:
		msg = "Missing files to search for TODOS comments" //todo review this? poor phrasing

	case NO_FILE_FOUND:
		msg = "No file found"

	case NO_TODO_COMMENT:
		msg = "No todo"

	default:
		msg = "Unknown error"
	}

	return msg + infos
}
