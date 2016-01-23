package main

import (
	"log"
	"encoding/json"
	"github.com/bradfitz/go-smtpd/smtpd"
)


type Envelope struct {
	To      string
	From    string
	Subject string
	Body    string
}

func (e *Envelope) AddRecipient(rcpt smtpd.MailAddress) error {
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
	data, err := json.Marshal(e)

	if err != nil {
		log.Println("ERROR:", err)
	}

	log.Println("MSG:", string(data))
	return nil
}

