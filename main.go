package goenv

import (
	"github.com/riccardooliva91/goenv/internal/bootstrap"
	"github.com/riccardooliva91/goenv/internal/env"
)

type Env env.Env

func Load(files []string, withOS bool) *env.Env {
	bootstrap.Init(files, withOS)

	return bootstrap.GetEnv()
}
