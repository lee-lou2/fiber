package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var MemoryCache = cache.New(
	5*time.Minute,
	10*time.Minute,
)
