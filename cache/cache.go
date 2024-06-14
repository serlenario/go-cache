package cache

type Cache struct {
	data map[string]interface{}
}

func New() *Cache {
	return &Cache{make(map[string]interface{})}
}

func (c *Cache) Set(key string, value interface{}) {
	c.data[key] = value
}

func (c *Cache) Get(key string) interface{} {
	if val, ok := c.data[key]; ok {
		return val
	}
	return nil
}

func (c *Cache) Delete(key string) {
	delete(c.data, key)
}
