package bootstrap

import (
	"github.com/riccardooliva91/goenv/internal/env"
	"github.com/riccardooliva91/goenv/internal/sources"
)

var variables env.Env
var shortCircuit bool = false

func Init(files []string, withOS bool) {
	if shortCircuit {
		return
	}

	variables = env.New(mergeSources(files, withOS))
	shortCircuit = true
}

func GetEnv() *env.Env {
	return &variables
}

func GetEnvCopy() env.Env {
	return variables
}

func mergeSources(files []string, withOS bool) map[string]string {
	res := map[string]string{}
	if withOS {
		res = sources.GetFromOs()
	}

	env := sources.GetFromEnv(files)
	for k, v := range env {
		res[k] = v
	}

	return res
}
