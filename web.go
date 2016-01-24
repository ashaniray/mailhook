package main

import (
	"fmt"
	"log"
	"net/http"
)

func StartWebInterface(host string, port int) {
	addr := fmt.Sprintf("%s:%d", host, port)
	log.Println("starting admin web interface on", addr)

	http.HandleFunc("/assets/", AssetHandler)
	http.HandleFunc("/", AdminHandler)
	http.HandleFunc("/new/", NewRuleHandler)
	http.HandleFunc("/create/", CreateRuleHandler)
	http.HandleFunc("/view/", ViewRuleHandler)
	http.HandleFunc("/edit/", EditRuleHandler)
	http.HandleFunc("/update/", UpdateRuleHandler)
	http.HandleFunc("/delete/", DeleteRuleHandler)

	err := http.ListenAndServe(addr, nil)

	if err != nil {
		log.Println("ERROR:", err)
	}
}
