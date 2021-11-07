package leetcode146

import (
	"fmt"
	"testing"
)

// hashtable
type LRUCache struct {
	cache map[int]int
	queue []int
	size  int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cache: make(map[int]int),
		queue: make([]int, capacity),
		size:  capacity,
	}
}

// first delete , join slice
func (this *LRUCache) Get(key int) int {
	_, ok := this.cache[key]
	if !ok {
		return -1
	}

	for i, v := range this.queue {
		if v == key {
			this.queue = append(this.queue[:i],this.queue[(i+1):]...)
			this.queue = append(this.queue,key)
		}
	}

	return this.cache[key]

}

// cache is exist ?
func (this *LRUCache) Put(key int, value int) {
	_, ok := this.cache[key]
	if !ok {
		// cache is full ?
		if this.size == len(this.queue) {
			delete(this.cache, this.queue[0])
			this.queue = this.queue[1:]
		}
		// join queue
		this.queue = append(this.queue, value)
	} else {
		// cache not exist
		// first delete , join slice
		for i, v := range this.queue {
			if v == key {
				this.queue = append(this.queue[:i],this.queue[(i+1):]...)
				this.queue = append(this.queue,value)
			}
		}
	}
	// join cache
	this.cache[key] = value
}

func TestLru(t *testing.T) {
	lruCache := Constructor(2)
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	lruCache.Get(1)
	lruCache.Put(3, 3)
	lruCache.Get(2)
	lruCache.Put(4, 4)
	lruCache.Get(1)
	lruCache.Get(3)
	lruCache.Get(4)
	fmt.Println(lruCache.queue)
}
