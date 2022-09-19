package env_test

import (
	"testing"

	"github.com/riccardooliva91/goenv/internal/env"
)

func TestGet(t *testing.T) {
	test := env.New(map[string]string{"test": "value"})

	exp1 := test.Get("test")
	exp2 := test.Get("test2")

	if exp1 != "value" {
		t.Fatalf("Expected test to be value, got %s instead", exp1)
	}
	if exp2 != "" {
		t.Fatalf("Expected test2 to be empty, got %s instead", exp2)
	}
}

func TestSet(t *testing.T) {
	test := env.New(map[string]string{})

	test.Set("test", "value")
	exp := test.Get("test")

	if exp != "value" {
		t.Fatalf("Expected test to be value, got %s instead", exp)
	}
}

func TestIsset(t *testing.T) {
	test := env.New(map[string]string{"test": "value"})

	exp1 := test.Isset("test")
	exp2 := test.Isset("test2")

	if !exp1 {
		t.Fatal("Expected test to be set, but it's not")
	}
	if exp2 {
		t.Fatal("Expected test2 to be unset, but it is")
	}
}

func TestDelete(t *testing.T) {
	test := env.New(map[string]string{"test": "value"})

	test.Delete("test")
	exp := test.Get("test")

	if exp != "" {
		t.Fatalf("Expected test to be deleted, but it's still there")
	}
}
