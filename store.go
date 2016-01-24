package main

import (
	"errors"
	"github.com/boltdb/bolt"
)

type Store struct {
	DB *bolt.DB
}

var DiskStore *Store

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
	err := s.DB.Update(func(tx *bolt.Tx) error {

		bucket, err := tx.CreateBucketIfNotExists(r.Bucket())
		if err != nil {
			return err
		}

		err = bucket.Put([]byte(r.Id), r.ToGob())

		if err != nil {
			return err
		}
		return nil
	})

	return err
}

func (s *Store) GetRule(id string) (*Rule, error) {
	return nil, nil
}

func (s *Store) GetAllRules() ([]*Rule, error) {
	ret := make([]*Rule, 0)
	s.DB.View(func(tx *bolt.Tx) error {

		bucket := tx.Bucket(RuleBucket())
		if bucket == nil {
			return errors.New("Bucket not found")
		}
		c := bucket.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			b := make([]byte, len(v))
			copy(b, v)
			r, err := NewRuleFromBytes(b)

			if err != nil {
				continue
			}
			ret = append(ret, r)
		}

		return nil
	})
	return ret, nil
}

func (s *Store) DeleteRule(id string) error {
	return nil
}
