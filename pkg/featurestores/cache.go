package featurestores

import "sync"

type Cache struct {
	inner FeatureStore
	cache map[string]bool
	lock  sync.Mutex
}

func NewCache(fs FeatureStore) FeatureStore {
	return &Cache{
		inner: fs,
		cache: make(map[string]bool),
		lock:  sync.Mutex{},
	}
}

func (c *Cache) Get(key string) (bool, error) {
	entry, ok := c.cache[key]
	if !ok {
		return c.tryAddKey(key)
	}

	return entry, nil
}

func (c *Cache) tryAddKey(key string) (bool, error) {
	entry, err := c.inner.Get(key)
	if err != nil {
		return false, err
	}

	c.lock.Lock()
	defer c.lock.Unlock()
	c.cache[key] = entry

	return entry, nil
}
