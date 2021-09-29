package cache

import (
	"errors"
	"github.com/dgraph-io/ristretto"
)

type MyCache struct {
	Cache *ristretto.Cache
}

var LocalStorage MyCache

func (Self *MyCache) Init() {
	var err error
	Self.Cache, err = ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		panic(err)
	}
}

func (Self *MyCache) SetKey(key, value interface{}) {
	// set a value with a cost of 1
	Self.Cache.Set(key, value, 1)

	// wait for value to pass through buffers
	Self.Cache.Wait()
}

func (Self *MyCache) GetKey(key interface{}) (value interface{}, err error) {
	value, found := Self.Cache.Get(key)
	if !found {
		// missing value
		return nil, errors.New("key not found")
	}

	return value, nil
}

func (Self *MyCache) DelKey(key interface{}) {
	Self.Cache.Del(key)
}
