package parser_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/riccardooliva91/goenv/internal/parser"
)

func TestParser(t *testing.T) {
	content, _ := os.ReadFile("../../tests/fixtures/.env3.fixture")

	exp := map[string]string{
		"TEST3":  "value3",
		"TEST4":  "value4",
		"TEST5":  "value5",
		"TEST6":  " val\"ue6",
		"TEST7":  "value7",
		"TEST8":  "value8",
		"TEST9":  "value9",
		"TEST10": "value10",
		"TEST11": " val\"ue11",
		"TEST12": "value12",
	}
	result := parser.ParseFileContent(content)

	if !reflect.DeepEqual(exp, result) {
		t.Fatalf("Parser failed, expected %s but got %s", exp, result)
	}
}

func TestParserPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected to panic, but didn't happen")
		}
	}()

	content, _ := os.ReadFile("../../tests/fixtures/.env4.fixture")
	parser.ParseFileContent(content)
}
