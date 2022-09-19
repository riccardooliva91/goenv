package merge_test

import (
	"bufio"
	"errors"
	"io"
	"io/fs"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/riccardooliva91/goenv/internal/merge"
)

type FakeFileInfo struct{}

func (f FakeFileInfo) Name() string       { return "" }
func (f FakeFileInfo) Size() int64        { return 1 }
func (f FakeFileInfo) Mode() os.FileMode  { return os.ModePerm }
func (f FakeFileInfo) ModTime() time.Time { return time.Time{} }
func (f FakeFileInfo) IsDir() bool        { return false }
func (f FakeFileInfo) Sys() any           { return nil }

func TestParseArgsNoFlags(t *testing.T) {
	args := []string{"file1", "file2"}
	merge.ParseArgs(args)

	f := merge.GetFiles()
	d := merge.GetDestination()
	o := merge.CanOverwrite()

	if !reflect.DeepEqual(args, f) {
		t.Fatal("Args do not match expectations")
	}
	if d != ".env" {
		t.Fatalf("Expected default .env, got %s instead", d)
	}
	if o {
		t.Fatal("Expected overwrite to be false, got true instead")
	}
}

func TestParseArgsWithFlags(t *testing.T) {
	cases := [][]string{
		{"file1", "file2", "-o", "-d", ".env.dest"},
		{"file1", "file2", "-d", ".env.dest", "-o"},
		{"file1", "file2", "--overwrite", "--destination", ".env.dest"},
	}
	for _, args := range cases {
		merge.ParseArgs(args)

		f := merge.GetFiles()
		d := merge.GetDestination()
		o := merge.CanOverwrite()

		if !reflect.DeepEqual(args[:2], f) {
			t.Fatal("Args do not match expectations")
		}
		if d != ".env.dest" {
			t.Fatalf("Expected default .env, got %s instead", d)
		}
		if !o {
			t.Fatal("Expected overwrite to be true, got false instead")
		}
	}
}

func TestParseArgsNoFiles(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("Expected test to panic, but it didn't")
		}
	}()

	args := []string{"file1", "file2", "-o", "-d"}
	merge.ParseArgs(args)
}

func TestParseArgsInvalidDestination(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("Expected test to panic, but it didn't")
		}
	}()

	args := []string{}
	merge.ParseArgs(args)
}

func TestOverwriteDoesNotExist(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("Expected test to not panic, but it did")
		}
	}()

	f := func(fn string) (fs.FileInfo, error) {
		return FakeFileInfo{}, os.ErrNotExist
	}
	args := []string{"file1", "file2"}
	merge.ParseArgs(args)
	merge.CheckWrite(f)
}

func TestOverwriteDoesExist(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("Expected test to panic, but it did't")
		}
	}()

	f := func(fn string) (fs.FileInfo, error) {
		return FakeFileInfo{}, os.ErrExist
	}
	args := []string{"file1", "file2"}
	merge.ParseArgs(args)
	merge.CheckWrite(f)
}

func TestOverwriteDoesntExistAndNoFlag(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("Expected test not to panic, but it did")
		}
	}()

	f := func(fn string) (fs.FileInfo, error) {
		return FakeFileInfo{}, nil
	}
	args := []string{"file1", "file2"}
	merge.ParseArgs(args)
	merge.CheckWrite(f)
}

func TestOverwrite(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("Expected test to not panic, but it did")
		}
	}()

	f := func(fn string) (fs.FileInfo, error) {
		return FakeFileInfo{}, os.ErrExist
	}
	args := []string{"file1", "file2", "-o"}
	merge.ParseArgs(args)
	merge.CheckWrite(f)
}

func TestWriteFailTouch(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("Expected test to panic, but it did't")
		}
	}()

	touch := func(name string, flags int, perm fs.FileMode) (*os.File, error) {
		return nil, errors.New("error")
	}
	w := func(iow io.Writer) *bufio.Writer {
		return bufio.NewWriter(os.Stdout)
	}
	env := map[string]string{"test": "val"}
	args := []string{"file1", "file2", "-o"}
	merge.ParseArgs(args)
	merge.Write(env, touch, w)
}

func TestWrite(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("Expected test to not panic, but it did")
		}
	}()

	touch := func(name string, flags int, perm fs.FileMode) (*os.File, error) {
		return nil, nil
	}
	w := func(iow io.Writer) *bufio.Writer {
		return bufio.NewWriter(os.Stdout)
	}
	env := map[string]string{"test": "val"}
	args := []string{"file1", "file2", "-o"}
	merge.ParseArgs(args)
	merge.Write(env, touch, w)
}
