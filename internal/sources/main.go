package sources

func GetFromOs() map[string]string {
	return getOsVars()
}

func GetFromEnv(files []string) map[string]string {
	return getEnvVars(files)
}
