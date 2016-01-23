package main

import (
	"flag"
)

var (
	smtpHost = flag.String("s", "0.0.0.0", "smtp bind address")
	smtpPort = flag.Int("p", 25, "smpt port")
)

func main() {
	flag.Parse()

	smtpOut := StartSMTPEndpoint(*smtpHost, *smtpPort)
	filterOut := StartFilter(smtpOut)
	StartDispatcher(filterOut)
}
