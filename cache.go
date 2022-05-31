package cache

import "time"

const (
	MinUint uint = 0 // binary: all zeroes

	// Perform a bitwise NOT to change every bit from 0 to 1
	MaxUint = ^MinUint // binary: all ones
)

type Cache struct {
	data map[string]item
}

type item struct {
	value    string
	deadline time.Time
}

func NewCache(key string, value string, deadline time.Time) Cache {
	return Cache{}
}

func (c *Cache) Get(key string) (string, bool) {
	//returns the value associated with the key and the boolean ok (true if exists, false if not), if the deadline of the key/value pair has not been exceeded yet
	val, ok := c.data[key]
	if ok && c.data[key].deadline.After(time.Now()) {
		return val.value, true
	} else {
		return "", false
	}
}

func (c *Cache) Put(key, value string) {
	//places a value with an associated key into cache. Value put with this method never expired(have infinite deadline). Putting into the existing key should overwrite the value
	c.data[key] = item{value, time.Unix(int64(MaxUint), 0)}
}

func (c *Cache) Keys() []string {
	//should return the slice of existing (non-expired keys)
	var res []string
	for k, _ := range c.data {
		if c.data[k].deadline.After(time.Now()) {
			res = append(res, k)
		}
	}
	return res
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.data[key] = item{value, deadline}
}
