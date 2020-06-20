package cache

import (
	"github.com/patrickmn/go-cache" //An in-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications.
	"sync"
	"time"
)

var (
	instance *cache.Cache
	once     sync.Once
)

//获取缓存单例
func Instance() *cache.Cache {
	once.Do(func() {
		c := cache.New(5*time.Minute, 10*time.Minute)
		instance = c
	})
	return instance
}
