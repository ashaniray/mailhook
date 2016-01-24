package main

import (
	"log"
)

func StartDispatcher(dispIn chan Packet) {
	log.Println("starting dispatcher")

	for {
		key := <-dispIn
		log.Println("DISPATCHER:", key)
	}
}
