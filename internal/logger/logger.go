package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

//log stuff in console or in a file ?
//requires time, guid etc ?

func WriteFileTest() {
	t := time.Now()

	logfilename := fmt.Sprintf("log_%d", t.Unix())
	fmt.Println(logfilename)

	s := "Un truc important"

	dirname := "logs"
	err := os.Mkdir(dirname, 0644)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filepath.Join(dirname, logfilename), []byte(s), 0644)
	if err != nil {
		panic(err)
	}
}
