package main

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"os"

	"github.com/riccardooliva91/goenv/internal/bootstrap"
	"github.com/riccardooliva91/goenv/internal/merge"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	merge.ParseArgs(os.Args[1:])
	merge.CheckWrite(getStatFunction())

	f := merge.GetFiles()
	bootstrap.Init(f, false)
	env := bootstrap.GetEnvCopy()
	merge.Write(env, getTouchFunction(), getWriterFunction())
}

func getStatFunction() merge.Stat {
	return func(fn string) (fs.FileInfo, error) {
		return os.Stat(fn)
	}
}

func getTouchFunction() merge.Touch {
	return func(name string, flags int, perm fs.FileMode) (*os.File, error) {
		return os.OpenFile(name, flags, perm)
	}
}

func getWriterFunction() merge.Writer {
	return func(f io.Writer) *bufio.Writer {
		return bufio.NewWriter(f)
	}
}
