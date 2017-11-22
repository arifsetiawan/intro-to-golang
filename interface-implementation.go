package main

// Memory storage
struct MemoryStorage {
	data map[string]string
}

func (m MemoryStorage) Create(key string, value string) {
	m.data[key] = value
}

// Redis storage
struct RedisStorage {
	client *redis.Client
}

func (m MemoryStorage) Create(key string, value string) {
	m.client.Set("key", "value", 0).Err()
}

// usage
var storage Storage
storage = new(RedisStorage)