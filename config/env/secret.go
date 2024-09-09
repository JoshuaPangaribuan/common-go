package env

func (e *envconfig) GetCredentials(key string) (string, string) {
	e.rwMutex.RLock()
	defer e.rwMutex.RUnlock()

	username := e.data[e.prefix+key+"_USERNAME"]
	password := e.data[e.prefix+key+"_PASSWORD"]

	return username, password
}

func (e *envconfig) GetAPIKey(key string) string {
	e.rwMutex.RLock()
	defer e.rwMutex.RUnlock()

	return e.data[e.prefix+key+"_API_KEY"]
}
