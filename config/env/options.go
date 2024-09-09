package env

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/joho/godotenv"
)

type EnvConfigOption func(*envconfig)

func WithPrefix(prefix string) EnvConfigOption {
	return func(e *envconfig) {
		e.prefix = prefix
	}
}

func WithSeparator(separator rune) EnvConfigOption {
	return func(e *envconfig) {
		e.separator = separator
	}
}

func WithFilenames(filenames ...string) EnvConfigOption {
	return func(e *envconfig) {
		e.filenames = filenames
	}
}

func WithWatcher() EnvConfigOption {
	return func(e *envconfig) {
		watchers, err := startWatchers(e)
		if err != nil {
			log.Println("error starting watchers:", err)
			return
		}
		e.watchers = watchers
	}
}

func startWatchers(e *envconfig) ([]*fsnotify.Watcher, error) {
	watchers := make([]*fsnotify.Watcher, 0, len(e.filenames))

	for _, filename := range e.filenames {
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			return nil, err
		}

		err = watcher.Add(filename)
		if err != nil {
			watcher.Close()
			return nil, err
		}

		go watchFileChanges(e, watcher)

		watchers = append(watchers, watcher)
	}

	return watchers, nil
}

func watchFileChanges(e *envconfig, watcher *fsnotify.Watcher) {
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Write == fsnotify.Write {
				data, err := godotenv.Read(e.filenames...)
				if err != nil {
					log.Println("error reading env file:", err)
					continue
				}
				e.rwMutex.Lock()
				e.data = data
				e.rwMutex.Unlock()
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("watcher error:", err)
		}
	}
}
