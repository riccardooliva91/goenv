package sources

import (
	"os"
	"strings"
)

var osVariables map[string]string

func getOsVars() map[string]string {
	osVariables = make(map[string]string)
	osVariables = getKeyValueMap(os.Environ())

	return osVariables
}

func getKeyValueMap(pairs []string) map[string]string {
	res := make(map[string]string)
	for _, str := range pairs {
		pair := strings.Split(str, "=")
		res[pair[0]] = pair[1]
	}

	return res
}
