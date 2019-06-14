package creates

import "sync"

type singleton map[string]string

var (
	instance singleton
	once sync.Once = sync.Once{}
)

func NewSingleton() singleton {
	once.Do(func() {
		instance = singleton{}
	})
	return instance
}

