package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetShouldShiftTheUsedItemToTheLastITemToTheRight(t *testing.T) {
	c := CreateCacheLRU(3)

	c.Set(1, 1)
	c.Set(2, 2)
	c.Set(3, 3)

	head := c.linkedList.Head.Value
	middle := c.linkedList.Head.Next.Value
	tail := c.linkedList.Tail.Value

	c.Get(2)

	newHead := c.linkedList.Head.Value
	newMiddle := c.linkedList.Head.Next.Value
	newTail := c.linkedList.Tail.Value

	// head 1 == 1
	assert.Equal(t, head, newHead)

	// middle 2 != 3
	assert.NotEqual(t, middle, newMiddle)
	// new middle 3 == 3
	assert.Equal(t, tail, newMiddle)

	// middle 2 == 2
	assert.Equal(t, middle, newTail)

}

func TestGetShouldReturnTheValueForAHitAndMinusOneForAMiss(t *testing.T) {
	scenarios := []struct {
		name     string
		cache    *CacheLRU
		key      int
		expected int
	}{
		{
			name:     "Return -1 for a MISS",
			cache:    CreateCacheLRU(2),
			key:      10,
			expected: -1,
		},
		{
			name: "Return the value for a HIT - 1 value",
			cache: func() *CacheLRU {
				c := CreateCacheLRU(2)
				c.Set(1, 1)
				return c
			}(),
			key:      1,
			expected: 1,
		},
		{
			name: "Return the value for a HIT - 3 values",
			cache: func() *CacheLRU {
				c := CreateCacheLRU(1)
				c.Set(1, 1)
				return c
			}(),
			key:      1,
			expected: 1,
		},
		{
			name: "Return the value for a HIT - 10 values",
			cache: func() *CacheLRU {
				c := CreateCacheLRU(100)
				c.Set(1, 1)
				c.Set(2, 2)
				c.Set(3, 3)
				c.Set(4, 4)
				c.Set(5, 5)
				c.Set(6, 6)
				c.Set(7, 7)
				c.Set(8, 8)
				c.Set(9, 9)
				c.Set(10, 10)
				return c
			}(),
			key:      5,
			expected: 5,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			result := scenario.cache.Get(scenario.key)

			assert.Equal(t, scenario.expected, result)
		})
	}
}

func TestSetShouldAddANewValueForTheCache(t *testing.T) {
	c := CreateCacheLRU(2)
	c.Set(1, 1)

	result := c.Get(1)

	assert.Equal(t, 1, result)
}

func TestSetShouldReplaceTheLRUItemWhenAtMaximunCapacity(t *testing.T) {
	c := CreateCacheLRU(2)
	c.Set(1, 1)
	c.Set(2, 2)

	c.Get(1)
	c.Set(3, 3)

	result := c.Get(2)

	assert.Equal(t, -1, result, "Not found because it was the least recently used item")
}
