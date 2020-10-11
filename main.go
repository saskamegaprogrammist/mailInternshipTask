package main

import (
	"github.com/saskamegaprogrammist/mailInternshipTask/requestParser"
	"log"
)

func main() {
	procNumber, errorsBool, err := requestParser.FlagsParse()
	if err != nil {
		log.Println(err)
		return
	}
	err = requestParser.ReadStdinWriteStdout(procNumber, errorsBool)
	if err != nil {
		log.Println(err)
	}
}
