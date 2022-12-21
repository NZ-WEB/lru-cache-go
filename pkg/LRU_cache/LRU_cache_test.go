package LRU_cache

import (
	assertInit "github.com/stretchr/testify/assert"
	"testing"
)

func TestLRUCache_NewLRUCache(t *testing.T) {
	assert := assertInit.New(t)
	instance := NewLRUCache(10)

	assert.Equal(10, instance.cap)

}

func TestLRUCache_Get(t *testing.T) {
	assert := assertInit.New(t)
	instance := NewLRUCache(3)

	instance.Add("1", "1")
	instance.Add("2", "2")

	assert.Equal(instance.Get("1"), "1")
	assert.Equal(instance.Get("2"), "2")
	assert.Equal(instance.Get("3"), "")
}

func TestLRUCache_Add(t *testing.T) {
	assert := assertInit.New(t)
	instance := NewLRUCache(3)
	TestKey := "1"
	TestVal := "1"

	assert.Equal(len(instance.cache), 0)
	instance.Add(TestKey, TestVal)
	assert.Equal(len(instance.cache), 1)
	assert.Equal(TestVal, instance.cache[TestKey])
	assert.Equal(instance.Add(TestKey, TestVal), false)

	instance.Add("2", "2")
	instance.Add("3", "3")

	assert.Equal(instance.cache["1"], "1")

	instance.Add("4", "4")
	assert.Equal(instance.cache["1"], "")
	assert.Equal(instance.cache["4"], "4")
}
