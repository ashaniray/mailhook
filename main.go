package main

import (
	"flag"
	"log"
)

var (
	smtpHost = flag.String("s", "0.0.0.0", "smtp server bind address.")
	smtpPort = flag.Int("p", 25, "smpt server port.")
	dbFile   = flag.String("d", "mailhook.db", "specify rules database file.")
	adminHost = flag.String("a", "0.0.0.0", "web server bind address.")
	adminPort = flag.Int("q", 8080, "webserver port.")
)

var MailStore = NewMemStore()

const defaultJs = `
func rule(data) {
	return true;
}
`

func DummyRules() {
	r1 := NewRule("rule1", defaultJs, []string{"http://ep1.com", "http://epr2.com"})
	r2 := NewRule("rule2", defaultJs, []string{"http://ep1.com", "http://epr2.com"})
	r3 := NewRule("rule3", defaultJs, []string{"http://ep1.com", "http://epr2.com"})
	r4 := NewRule("rule4", defaultJs, []string{"http://ep1.com", "http://epr2.com"})

	DiskStore.SaveRule(r1)
	DiskStore.SaveRule(r2)
	DiskStore.SaveRule(r3)
	DiskStore.SaveRule(r4)

	rxs, _ := DiskStore.GetAllRules()

	for _, rx := range rxs {
		log.Println("RULE:", rx)
	}

}

func main() {
	flag.Parse()

	var err error
	DiskStore, err = NewStore(*dbFile)
	defer DiskStore.Close()

	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	smtpOut := StartSMTPEndpoint(*smtpHost, *smtpPort)
	filterOut := StartFilter(smtpOut)
	go StartDispatcher(filterOut)

	StartWebInterface(*adminHost, *adminPort)
}
