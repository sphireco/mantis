package cache

import (
	"github.com/allegro/bigcache"
	mantisError "github.com/sphireco/mantis/error"
)

type BigCache struct {
	Cache  *bigcache.BigCache
	Config bigcache.Config
}

// Init
func (b *BigCache) Init() error {
	var err error
	b.Cache, err = bigcache.NewBigCache(b.Config)

	if err != nil {
		mantisError.HandleFatalError(err)
		return err
	}

	return nil
}
