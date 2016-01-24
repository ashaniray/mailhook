package main

import (
	"log"
)

func StartFilter(in chan string) chan string {
	filterOut := make(chan string)
	log.Println("starting filter ...")

	go func() {
		for {
			key := <-in
			log.Println("FILTER:", key)
			filterOut <- key
		}
	}()

	return filterOut
}
