package commongo

type Config interface {
	GetObject(key string) any
	GetString(key string) string
	GetInt(key string) int64
	GetFloat(key string) float64
	GetUint(key string) uint64
	GetBool(key string) bool
	GetArray(key string) []any
	GetMap(key string) map[string]any
}

type Secret interface {
	GetCredentials(key string) (string, string)
	GetAPIKey(key string) string
}

var _ error = (*ConfigError)(nil)

type ConfigError struct {
	Message string
}

func (e *ConfigError) Error() string {
	return e.Message
}
