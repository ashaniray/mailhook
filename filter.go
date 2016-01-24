package main

import (
	"encoding/json"
	"log"
)

type Packet struct {
	Key       string
	Endpoints []string
}

func NewPacket(key string, eps []string) *Packet {
	return &Packet{Key: key, Endpoints: eps}
}

func lookupMessage(key string) *Message {
	mail := MailStore.Get(key)
	message := new(Message)
	json.Unmarshal([]byte(mail), message)
	return message

}

func applyRule(rule *Rule, key string, fout chan Packet) {
	msg := lookupMessage(key)
	if rule.Evaluate(msg) {
		fout <- *NewPacket(key, rule.Endpoints)
	}
}

func applyRules(key string, fout chan Packet) {
	rules, err := DiskStore.GetAllRules()

	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	for _, rule := range rules {
		go applyRule(rule, key, fout)
	}

}

func StartFilter(in chan string) chan Packet {
	filterOut := make(chan Packet)
	log.Println("starting filter ...")

	go func() {
		for {
			key := <-in
			log.Println("FILTER:", key)
			applyRules(key, filterOut)
		}
	}()

	return filterOut
}
