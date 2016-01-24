package main

import (
	"encoding/json"
	"github.com/bradfitz/go-smtpd/smtpd"
	"log"
	"strings"
)

type Message struct {
	From string
	To   []string
	Body string
}

type Envelope struct {
	*smtpd.BasicEnvelope
	msg Message
}

func (e *Envelope) AddRecipient(rcpt smtpd.MailAddress) error {
	e.msg.To = append(e.msg.To, rcpt.Email())
	return e.BasicEnvelope.AddRecipient(rcpt)
}

func (e *Envelope) Write(line []byte) error {
	str := strings.Replace(string(line), "\n", " ", -1)
	str = strings.Replace(str, "\r", " ", -1)
	e.msg.Body += str
	return nil
}

func (e *Envelope) Close() error {
	messageInJson, err := json.Marshal(e.msg)
	if err != nil {
		log.Printf("Error occured while converting to json", err)
		return err
	}
	key := MailStore.Save(string(messageInJson))
	log.Printf("The message received : %q", string(messageInJson))
	globalOut <- key
	return nil
}
