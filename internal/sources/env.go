package sources

import (
	"log"
	"os"

	"github.com/riccardooliva91/goenv/internal/parser"
)

var envVariables map[string]string
var envFileName string = ".env"

func getEnvVars(files []string) map[string]string {
	warn := "WARNING: env file \"%s\" not found!"
	sources := []string{envFileName}
	if len(files) > 0 {
		sources = files
	}

	envVariables = make(map[string]string)
	for _, file := range sources {
		content, err := os.ReadFile(file)
		if err != nil {
			log.Printf(warn, file)

			continue
		}

		parsed := parser.ParseFileContent(content)
		for k, v := range parsed {
			envVariables[k] = v
		}
	}

	return envVariables
}
