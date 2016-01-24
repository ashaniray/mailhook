package main

import (
	"github.com/boltdb/bolt"
)

type Store struct {
	DB *bolt.DB
}

func NewStore(name string) (*Store, error) {
	db, err := bolt.Open(name, 0600, nil)

	if err != nil {
		return nil, err
	}

	return &Store{DB: db}, nil
}

func (s *Store) Close() {
	s.DB.Close()
}

func (s *Store) SaveRule(r *Rule) error {
	return nil
}

func (s *Store) GetRule(id string) (*Rule, error) {
	return nil, nil
}

func (s *Store) GetAllRules() ([]*Rule, error) {
	return nil, nil
}

func (s *Store) DeleteRule(id string) error {
	return nil
}
