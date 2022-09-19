package bootstrap_test

import (
	"reflect"
	"testing"

	"github.com/riccardooliva91/goenv/internal/bootstrap"
)

func TestInit(t *testing.T) {
	files := []string{"../../tests/fixtures/.env1.fixture", "../../tests/fixtures/.env2.fixture"}
	bootstrap.Init(files, true)
	env1 := bootstrap.GetEnv()
	bootstrap.Init([]string{}, true)
	env2 := bootstrap.GetEnv()
	bootstrap.Init([]string{}, true)
	env3 := bootstrap.GetEnvCopy()

	if !reflect.DeepEqual(env1, env2) || reflect.DeepEqual(env1, env3) {
		t.Fatal("Expected initialization to happen once, was initialized twice")
	}
}

func TestMergeSources(t *testing.T) {
	files := []string{"../../tests/fixtures/.env1.fixture", "../../tests/fixtures/.env2.fixture", "donotexist"}
	bootstrap.Init(files, true)
	env := bootstrap.GetEnv()

	test1 := env.Get("TEST1")
	test2 := env.Get("TEST2")
	overridden := env.Get("TOBEOVERRIDDEN")
	if test1 != "value1" {
		t.Fatalf("Expected TEST1 to be value1, got %s instead", test1)
	}
	if test2 != "value2" {
		t.Fatalf("Expected TEST2 to be value2, got %s instead", test2)
	}
	if overridden != "overridden" {
		t.Fatalf("Expected variable to be overridden, got %s instead", overridden)
	}
}
