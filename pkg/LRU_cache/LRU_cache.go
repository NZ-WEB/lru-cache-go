package LRU_cache

import (
	"container/list"
	"log"
)

type LRUCache struct {
	cache map[string]string
	cap   int
	queue *list.List
}

type ListItem struct {
	Key string
	Val string
}

func NewLRUCache(n int) *LRUCache {
	c := &LRUCache{
		cache: make(map[string]string, n),
	}
	c.cap = n
	c.queue = list.New()

	return c
}

func (c *LRUCache) purge() {
	if element := c.queue.Back(); element != nil {
		item := c.queue.Remove(element).(*ListItem)
		delete(c.cache, item.Key)
	}
}

func (c *LRUCache) Add(key, value string) bool {
	if c.cache[key] == value {
		return false
	}

	l := len(c.cache)
	ca := c.cap

	log.Println(l, ca)

	if len(c.cache) == c.cap {
		c.purge()
	}

	item := &ListItem{
		key,
		value,
	}

	c.queue.PushFront(item)
	c.cache[item.Key] = item.Val

	return true
}

func (c *LRUCache) Get(key string) string {
	return c.cache[key]
}
