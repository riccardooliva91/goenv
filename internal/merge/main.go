package merge

import (
	"fmt"
	"strings"
)

var files []string
var destination string
var overwrite bool

func ParseArgs(args []string) {
	help := "USAGE: goenv-merge <file_1>[,file_2 ...] [arguments [-d|--destination dest] [-o|--overwrite]]"
	f, d, o, valid := validate(args)
	if !valid {
		panic(help)
	}

	files = f
	destination = d
	overwrite = o
}

func CanOverwrite() bool {
	return overwrite
}

func GetDestination() string {
	return destination
}

func CheckWrite(s Stat) {
	exists := fileExists(s)
	if exists && !overwrite {
		panic(fmt.Sprintf("file %s cannot be overridden, flag missing", destination))
	}
}

func Write(e map[string]string, t Touch, w Writer) {
	f, err := createFile(t)
	if err != nil {
		panic(err)
	}

	writeEnv(e, f, w)
}

func GetFiles() []string {
	return files
}

func getContent(e map[string]string) []byte {
	lines := []string{}
	for k, v := range e {
		lines = append(lines, fmt.Sprintf("%s=\"%s\"", k, v))
	}
	str := strings.Join(lines, "\n")

	return []byte(str)
}
