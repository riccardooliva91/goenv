package parser

import (
	"fmt"
	"regexp"
	"strings"
)

func ParseFileContent(rawContent []byte) map[string]string {
	var separator string

	content := string(rawContent)
	parsed := make(map[string]string)
	lines := strings.Split(content, "\n")
	for i := range lines {
		line := removeComments(lines[i])
		if len(line) == 0 {
			continue
		}

		equalIndex := strings.Index(lines[i], "=")
		columnIndex := strings.Index(lines[i], ":")
		isEqual := equalIndex > -1 && (equalIndex < columnIndex || columnIndex == -1)
		isColumn := columnIndex > -1 && (columnIndex < equalIndex || equalIndex == -1)
		if isEqual {
			separator = "="
		} else if isColumn {
			separator = ":"
		} else {
			panic(fmt.Errorf("invalid entry: %s", lines[i]))
		}

		pair := strings.Split(line, separator)
		parsed[cleanString(pair[0])] = cleanString(pair[1])
	}

	return parsed
}

func removeComments(line string) string {
	re := regexp.MustCompile(`#.*`)

	return re.ReplaceAllString(line, "")
}

func cleanString(str string) string {
	res := strings.Trim(str, " ") // spaces first, so we keep the ones inside the double quotes
	res = strings.Trim(res, "\"")

	return res
}
