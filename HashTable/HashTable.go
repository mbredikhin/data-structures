package HashTable

import (
	"errors"
	"hash/fnv"
)

// HashTable data structure
type HashTable struct {
	memory []int
}

// NewHashTable - HashTable constructor
func NewHashTable() *HashTable {
	var memory []int
	return &HashTable{memory}
}

// Set value by key
func (h *HashTable) Set(key string, value int) {
	hash := generateHash(key)
	if uint8(len(h.memory)) < hash {
		m := make([]int, hash+1)
		copy(m, h.memory)
		h.memory = m
	}
	h.memory[hash] = value
}

// Get value by key
func (h *HashTable) Get(key string) (int, error) {
	hash := generateHash(key)
	if uint8(len(h.memory)-1) < hash {
		return 0, errors.New("Invalid key")
	}
	return h.memory[hash], nil
}

// Remove value by key
func (h *HashTable) Remove(key string) error {
	hash := generateHash(key)
	if uint8(len(h.memory)-1) < hash {
		return errors.New("Invalid key")
	}
	h.memory = append(h.memory[:hash], h.memory[hash+1:]...)
	return nil
}

func generateHash(s string) uint8 {
	hash := fnv.New32a()
	hash.Write([]byte(s))
	return uint8(hash.Sum32() % 256)
}
