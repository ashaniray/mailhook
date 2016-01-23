package main

import (
	"log"
	"github.com/bradfitz/go-smtpd/smtpd"
)


type Envelope struct {
	To      string
	From    string
	Subject string
	Body    string
}

func (e *Envelope) AddRecipient(rcpt smtpd.MailAddress) error {
	log.Println("AddRecipient")
	e.To = rcpt.Email()
	return nil
}

func  (e *Envelope) BeginData() error {
	return nil
}

func  (e *Envelope) Write(line []byte) error {
	e.Body += string(line)
	return nil
}

func  (e *Envelope) Close() error {
	log.Println("Close", e)
	return nil
}

