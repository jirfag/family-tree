package db

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var CacheClient = cacheClient()

func cacheClient() (c *cache.Cache) {
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	c = cache.New(cache.NoExpiration, 10*time.Minute)
	return c
}
