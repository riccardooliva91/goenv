package merge

func validate(args []string) ([]string, string, bool, bool) {
	files, params := splitFilesAndFlags(args)
	if len(files) == 0 {
		return []string{}, "", false, false
	}

	destination, overwrite, valid := parseParams(params)

	return files, destination, overwrite, valid
}

func splitFilesAndFlags(args []string) ([]string, []string) {
	files := []string{}
	params := []string{}
	for i, v := range args {
		if isOverWriteFlag(v) || isDestinationFlag(v) {
			files = args[:i]
			params = args[i:]
			break
		}
	}
	if len(params) == 0 {
		files = args
	}

	return files, params
}

func parseParams(params []string) (string, bool, bool) {
	destination := ".env"
	overwrite := false
	valid := true

	for i, v := range params {
		if isOverWriteFlag(v) {
			overwrite = true
		} else if isDestinationFlag(v) {
			if len(params) == i+1 {
				valid = false
				continue // Has destination not specified
			} else if len(params) >= i+2 && !isDestinationFlag(params[i+1]) && !isOverWriteFlag(params[i+1]) {
				destination = params[i+1]
			}
		}
	}

	return destination, overwrite, valid
}

func isOverWriteFlag(v string) bool {
	return v == "-o" || v == "--overwrite"
}

func isDestinationFlag(v string) bool {
	return v == "-d" || v == "--destination"
}
