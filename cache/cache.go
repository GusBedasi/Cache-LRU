package cache

type cache interface {
	Get(key int) int
	Set(key, value int)
}

type CacheLRU struct {
	used       int
	capacity   int
	linkedList linkedList
	hashMap    map[int]*node
}

func CreateCacheLRU(capacity int) *CacheLRU {
	return &CacheLRU{
		capacity: capacity,
		used:     0,
		hashMap:  make(map[int]*node, capacity),
	}
}

func (c *CacheLRU) Get(key int) int {
	if n := c.hashMap[key]; n != nil {
		value := n.Value
		c.linkedList.remove(n)
		c.linkedList.insert(n)
		return value
	} else {
		return -1
	}
}

func (c *CacheLRU) Set(key, value int) {
	if c.used == c.capacity {
		//Least recently used
		lru := c.linkedList.Head

		//Remove from the hashmap and linked list
		c.linkedList.remove(lru)
		delete(c.hashMap, lru.Key)

		//Create new node
		n := &node{Key: key, Value: value}

		//Adds to hashmap and linked list
		c.hashMap[key] = n
		c.linkedList.insert(n)
	} else {
		//Create new node
		n := &node{Key: key, Value: value}

		//Adds to hashmap and linked list
		c.hashMap[key] = n
		c.linkedList.insert(n)

		c.used += 1
	}
}
