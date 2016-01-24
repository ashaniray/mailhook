package main

import (
	"bytes"
	"encoding/gob"
	"github.com/pborman/uuid"
	"github.com/robertkrimen/otto"
	"log"
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

func (r *Rule) Evaluate(payload *Message) bool {
	result, err := evaluateRule(r.Script, payload)
	if err != nil {
		log.Println("ERROR:", err)
		return false
	}

	return result
}

func evaluateRule(src string, payload *Message) (bool, error) {
	js := otto.New()
	var ruleFunc otto.Value
	js.Set("rule", func(call otto.FunctionCall) otto.Value {
		ruleFunc = call.Argument(0)
		return otto.UndefinedValue()
	})

	js.Run(src)
	arg, err := js.ToValue(payload)

	if err != nil {
		return false, err
	}

	ret, err := ruleFunc.Call(otto.NullValue(), arg)

	if err != nil {
		return false, err
	}

	return ret.ToBoolean()
}
