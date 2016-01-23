package main

import (
	"log"
)

func StartFilter(in chan string) chan string {
	filterOut := make(chan string)
	log.Println("starting filter ...")

	return filterOut
}
