package merge

import (
	"bufio"
	"errors"
	"io"
	"io/fs"
	"os"
)

type Stat func(string) (fs.FileInfo, error)
type Touch func(string, int, fs.FileMode) (*os.File, error)
type Writer func(io.Writer) *bufio.Writer

func fileExists(s Stat) bool {
	var exists bool
	if _, err := s(destination); err == nil {
		exists = true
	} else if errors.Is(err, os.ErrNotExist) {
		exists = false
	} else {
		exists = false
	}

	return exists
}

func createFile(t Touch) (*os.File, error) {
	return t(destination, os.O_RDWR|os.O_CREATE, 0644)
}

func writeEnv(env map[string]string, file *os.File, wr Writer) {
	content := getContent(env)
	w := wr(file)
	_, err := w.Write(content)
	if err != nil {
		panic(err)
	}

	w.Flush()
}
