package go_cacher

type Cache struct {
	storage map[string]any
}

func New() *Cache {
	return &Cache{storage: make(map[string]any)}
}

func (c *Cache) Set(key string, value any) {
	c.storage[key] = value
}

func (c *Cache) Get(key string) any {
	value, exists := c.storage[key]

	if exists {
		return value
	}

	return nil
}

func (c *Cache) Delete(key string) any {
	_, exists := c.storage[key]

	if exists {
		delete(c.storage, key)
	}

	return nil
}
