package main

import (
	"sync"
	"fmt"
	"crypto/md5"
)

type MemStore struct {
	Storage map[string]string
	Guard sync.Mutex
}

func NewMemStore() *MemStore {
	s := new(MemStore)
	s.Storage = make(map[string]string)
	return s
}

func (s *MemStore) Save(content string)  string {
	s.Guard.Lock()
	defer s.Guard.Unlock()

	key := checksum(content)
	s.Storage[key] = content

	return key
}


func (s *MemStore) Get(key string) string {
	data, ok := s.Storage[key]

	if ok {
		return data
	} else {
		return ""
	}
}

func checksum(data string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}
