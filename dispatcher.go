package main

import (
	"log"
	"net/http"
	"strings"
)

func dispatchPayload(ep string, payload string) {
	resp, err := http.Post(ep, "application/json", strings.NewReader(payload))
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(resp)
}

func StartDispatcher(dispIn chan Packet) {
	log.Println("starting dispatcher")

	for {
		packet := <-dispIn
		payload := MailStore.Get(packet.Key)

		log.Println("dispatching: message", packet.Key, "to", packet.Endpoints)

		for _, ep := range packet.Endpoints {
			go dispatchPayload(ep, payload)
		}
	}
}
