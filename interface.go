package main

type Storage interface {
	Put(key string, data string) error
	Get(key string) (string, error)
	Delete(key string) error
}
