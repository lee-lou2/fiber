package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var MemoryCache = cache.New(
	5*time.Minute,
	10*time.Minute,
)
