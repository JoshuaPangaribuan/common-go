package env

import (
	"os"
	"strconv"
	"strings"
)

func (e *envconfig) GetObject(key string) any {
	e.rwMutex.RLock()
	defer e.rwMutex.RUnlock()
	return os.Getenv(e.prefix + key)
}

func (e *envconfig) GetString(key string) string {
	e.rwMutex.RLock()
	defer e.rwMutex.RUnlock()
	return os.Getenv(e.prefix + key)
}

func (e *envconfig) GetInt(key string) int64 {
	e.rwMutex.RLock()
	defer e.rwMutex.RUnlock()

	intRes, err := strconv.ParseInt(os.Getenv(e.prefix+key), 10, 64)
	if err != nil {
		intRes = 0
	}

	return intRes
}

func (e *envconfig) GetFloat(key string) float64 {
	e.rwMutex.RLock()
	defer e.rwMutex.RUnlock()

	floatRes, err := strconv.ParseFloat(os.Getenv(e.prefix+key), 64)
	if err != nil {
		floatRes = 0
	}

	return floatRes
}

func (e *envconfig) GetUint(key string) uint64 {
	e.rwMutex.RLock()
	defer e.rwMutex.RUnlock()

	uintRes, err := strconv.ParseUint(os.Getenv(e.prefix+key), 10, 64)
	if err != nil {
		uintRes = 0
	}

	return uintRes
}

func (e *envconfig) GetBool(key string) bool {
	e.rwMutex.RLock()
	defer e.rwMutex.RUnlock()

	boolRes, err := strconv.ParseBool(os.Getenv(e.prefix + key))
	if err != nil {
		boolRes = false
	}

	return boolRes
}

func (e *envconfig) GetArray(key string) []any {
	e.rwMutex.RLock()
	defer e.rwMutex.RUnlock()

	strArr := strings.Split(os.Getenv(e.prefix+key), ",")
	result := make([]any, len(strArr))

	for i, v := range strArr {
		result[i] = v
	}

	return result
}

func (e *envconfig) GetMap(key string) map[string]any {
	e.rwMutex.RLock()
	defer e.rwMutex.RUnlock()

	strArr := strings.Split(os.Getenv(e.prefix+key), ",")
	result := make(map[string]any)

	for _, v := range strArr {
		keyVal := strings.Split(v, ":")
		if len(keyVal) == 2 {
			result[keyVal[0]] = keyVal[1]
		}
	}

	return result
}
