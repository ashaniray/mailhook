package main

import (
	"log"
)

func StartDispatcher(dispIn chan string) {
	log.Println("starting dispatcher")

	data := <- dispIn

	log.Println(data)
}
