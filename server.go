package main

import (
	"log"
	"fmt"
	"github.com/bradfitz/go-smtpd/smtpd"
)

var chIn = make(chan string)

func newMail(c smtpd.Connection, from smtpd.MailAddress) (smtpd.Envelope, error) {
	log.Println("NEW MAIL", from)
	return new(Envelope), nil
}

func startServer(host string, port int) {
	addr := fmt.Sprintf("%s:%d", host, port)
	log.Println("Starting smtp endpoint on", addr)

	s := &smtpd.Server{Addr: addr, OnNewMail: newMail }
	err := s.ListenAndServe()

	if err != nil {
		log.Println("ERROR:", err)
	}

}


func StartSMTPEndpoint(host string, port int) chan string {
	smtpOut := make(chan string)

	go startServer(host, port)
	go func() {
		key := <-chIn
		smtpOut <- key
	}()

	return smtpOut
}
