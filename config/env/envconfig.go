package env

import (
	"io"
	"log"
	"os"
	"sync"

	common "github.com/JoshuaPangaribuan/common-go"
	"github.com/fsnotify/fsnotify"
	"github.com/joho/godotenv"
)

var _ common.Secret = (*envconfig)(nil)
var _ common.Config = (*envconfig)(nil)

type envconfig struct {
	io.Closer
	prefix    string
	separator rune
	filenames []string
	rwMutex   sync.RWMutex
	data      map[string]string
	watchers  []*fsnotify.Watcher
}

func NewConfig(opts ...EnvConfigOption) (*envconfig, error) {
	conf := defaults()

	// First load config
	data, err := godotenv.Read(conf.filenames...)
	if err != nil {
		return nil, err
	}

	conf.data = data

	// Then load options
	for _, opt := range opts {
		opt(conf)
	}

	return conf, nil
}

func defaults() *envconfig {
	return &envconfig{
		prefix:    "",
		separator: '_',
		filenames: []string{".env"},
	}
}

func (e *envconfig) Close() error {
	e.rwMutex.Lock()
	defer e.rwMutex.Unlock()

	// Close all watchers
	for _, watcher := range e.watchers {
		if err := watcher.Close(); err != nil {
			log.Println("error closing watcher:", err)
		}
	}

	// Clear environment variables
	os.Clearenv()

	return nil
}
