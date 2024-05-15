package postgres

import (
	"fmt"
	"sync"

	"github.com/KaffeeMaschina/http-rest-api/internal/storage"
)

type Cache struct {
	m      sync.RWMutex
	orders map[string]storage.Orders
}

// Create new cache
func NewCashe() *Cache {
	orders := make(map[string]storage.Orders)

	return &Cache{orders: orders}
}

// Write data to cache
func (c *Cache) SetCache(oid string, o storage.Orders) {
	c.m.Lock()
	defer c.m.Unlock()

	c.orders[oid] = o
	fmt.Println(o)

}

// Read data from cache
func (c *Cache) OrderOutCache(oid string) (o storage.Orders) {

	c.m.RLock()

	o = c.orders[oid]
	c.m.RUnlock()

	return o

}
