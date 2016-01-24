package main

import (
	"github.com/pborman/uuid"
)

type Rule struct {
	Id string
	Title string
	Script string
	Endpoints []string
}

func NewRule(title string, src string, eps []string) *Rule {
	return &Rule {
		Id: uuid.NewUUID().String(),
		Title: title,
		Script: src,
		Endpoints: eps}
}
