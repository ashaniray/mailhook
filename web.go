package main

import (
	"log"
	"net/http"
	"fmt"
)

func AdminHandler(w http.ResponseWriter, r *http.Request) { 
	w.Write([]byte(`
<html><h1>mailhook</h1></html>
	`))
}


func StartWebInterface(host string, port int) {
	addr := fmt.Sprintf("%s:%d", host, port)
	log.Println("starting admin web interface on", addr) 

	http.HandleFunc("/", AdminHandler)
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		log.Println("ERROR:",err)
	}
}
