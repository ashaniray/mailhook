package main

import (
	"log"
	"fmt"
)


func StartSMTPEndpoint(host string, port int) chan string {
	addr := fmt.Sprintf("%s:%d", host, port)
	log.Println("Starting smtp endpoint on", addr)

	smtpOut := make(chan string)

	return smtpOut
}
