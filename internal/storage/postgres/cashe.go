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

func NewCashe() *Cache {
	orders := make(map[string]storage.Orders)

	return &Cache{orders: orders}
}

func (c *Cache) SetCache(oid string, o storage.Orders) {
	c.m.Lock()
	defer c.m.Unlock()

	c.orders[oid] = o

	fmt.Println(c.orders)
}

func (c *Cache) GetCacheFromDB() {

}
