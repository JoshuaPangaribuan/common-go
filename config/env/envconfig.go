package env

import "sync"

type envconfig struct {
	prefix    string
	separator rune
	filenames []string
	rwMutex   sync.RWMutex
}

func NewConfig() *envconfig {
	return defaults()
}

func defaults() *envconfig {
	return &envconfig{
		prefix:    "",
		separator: '_',
		filenames: []string{".env"},
	}
}
