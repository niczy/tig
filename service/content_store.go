package service

import (
	"crypto/sha256"
	"fmt"
)

type ContentStore struct {
	kv map[string]string
}

var contentStore = &ContentStore{
	kv: make(map[string]string),
}

// Store data, returns the hash key of the data.
func (c *ContentStore) Put(data string) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", data)))

	hash := fmt.Sprintf("%x", h.Sum(nil))
	c.kv[hash] = data
	return hash
}

func (c *ContentStore) Get(contentHash string) string {
	return c.kv[contentHash]
}
