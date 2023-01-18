package cache

import (
	"time"

	"github.com/bluele/gcache"

	"github.com/starudream/go-lib/config"
)

type cacheType string

const (
	TypeSIMPLE cacheType = gcache.TYPE_SIMPLE
	TypeLRU    cacheType = gcache.TYPE_LRU
	TypeLFU    cacheType = gcache.TYPE_LFU
	TypeARC    cacheType = gcache.TYPE_ARC
)

var (
	size   int
	expire time.Duration
)

func init() {
	size, expire = config.GetInt("cache.size"), config.GetDuration("cache.expire")
	if size == 0 {
		size = 128
	}
	if expire == 0 {
		expire = 24 * time.Hour
	}
}

func New(size int, evictType cacheType, expire time.Duration) gcache.Cache {
	return gcache.New(size).Expiration(expire).EvictType(string(evictType)).Build()
}

// SIMPLE no clear priority for evict cache. It depends on key-value map order.
func SIMPLE() gcache.Cache {
	return New(size, TypeSIMPLE, expire)
}

// LRU discards the least recently used items first.
func LRU() gcache.Cache {
	return New(size, TypeLRU, expire)
}

// LFU discards the least frequently used items first.
func LFU() gcache.Cache {
	return New(size, TypeLFU, expire)
}

// ARC constantly balances between LRU and LFU, to improve the combined result.
func ARC() gcache.Cache {
	return New(size, TypeARC, expire)
}
