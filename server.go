package main

import (
	"github.com/bradfitz/go-smtpd/smtpd"
	"log"
	"strconv"
)

var myMessage Message

var globalOut = make(chan string)

func StartSMTPEndpoint(host string, port int) chan string {
	smtpOut := make(chan string)

	go func() {
		for {
			out := <-globalOut
			smtpOut <- out
		}
	}()

	go func() {
		s := &smtpd.Server{
			Addr:      ":" + strconv.Itoa(port),
			Hostname:  host,
			OnNewMail: onNewMail,
		}

		// entry point to the SMTP endpoint
		err := s.ListenAndServe()
		if err != nil {
			log.Fatalf("ListenAndServe: %v", err)
		}

	}()

	return smtpOut
}

func onNewMail(c smtpd.Connection, from smtpd.MailAddress) (smtpd.Envelope, error) {
	log.Printf("New mail received from %q", from)
	myMessage.From = from.Email()
	return &Envelope{new(smtpd.BasicEnvelope), myMessage}, nil
}
