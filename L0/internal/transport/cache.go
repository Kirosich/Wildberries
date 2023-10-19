package transport

import (
	"Wildber/internal/database"
	"Wildber/internal/models"
	"fmt"
	"sync"
)

type Cache struct {
	mu    sync.RWMutex
	store map[string]models.Order
}

func (c *Cache) Set(key_uid string, value_model models.Order) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key_uid] = value_model
}

func (c *Cache) Get(key_uid string) (models.Order, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value_model, ok := c.store[key_uid]
	return value_model, ok
}

func LoadCache(db *database.Database) (*Cache, error) {
	c := &Cache{
		store: make(map[string]models.Order),
	}
	uid, information, err := db.SelectAll()
	if err != nil {
		fmt.Println("Error while Loading Cash")
		return nil, err
	}
	for i := 0; i < len(uid); i++ {
		c.Set(uid[i], information[i])
	}
	return c, nil
}
