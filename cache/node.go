package cache

type node struct {
	Key, Value int
	Prev       *node
	Next       *node
}
