package main

import (
	"github.com/pborman/uuid"
	"bytes"
	"encoding/gob"
)

type Rule struct {
	Id        string
	Title     string
	Script    string
	Endpoints []string
}

func NewRule(title string, src string, eps []string) *Rule {
	return &Rule{
		Id:        uuid.NewUUID().String(),
		Title:     title,
		Script:    src,
		Endpoints: eps}
}

func NewRuleFromBytes(b []byte) (*Rule, error) {
	buf := bytes.NewBuffer(b)
	var r Rule
	dec := gob.NewDecoder(buf)
	err := dec.Decode(&r)
	return &r, err
}

func RuleBucket() []byte {
	return []byte("rules")
}

func (r *Rule) Bucket() []byte {
	return RuleBucket()
}

func (r *Rule) ToGob() []byte {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)

	err := enc.Encode(r)

	if err != nil {
		return nil
	}

	return buff.Bytes()
}
