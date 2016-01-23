package main

import (
	"log"
	"fmt"
	"github.com/bradfitz/go-smtpd/smtpd"
)

type Envelope struct {
	Rcpts []string
	Message string
}

func (e *Envelope) AddRecipient(rcpt smtpd.MailAddress) error {
	log.Println("AddRecipient")
	return nil
}

func  (e *Envelope) BeginData() error {
	log.Println("BeginData")
	return nil
}

func  (e *Envelope) Write(line []byte) error {
	log.Println("Write")
	return nil
}

func  (e *Envelope) Close() error {
	log.Println("Close")
	return nil
}

func newMail(c smtpd.Connection, from smtpd.MailAddress) (smtpd.Envelope, error) {
	log.Println("NEW MAIL", from)
	return new(Envelope), nil
}



func StartSMTPEndpoint(host string, port int) chan string {
	smtpOut := make(chan string)

	go func() {
		addr := fmt.Sprintf("%s:%d", host, port)
		log.Println("Starting smtp endpoint on", addr)

		s := &smtpd.Server{Addr: addr, OnNewMail: newMail }
		err := s.ListenAndServe()

		if err != nil {
			log.Println("ERROR:", err)
		}
	}()

	return smtpOut
}
