package main

import (
	"flag"
	"github.com/saskamegaprogrammist/mailInternshipTask/requestParser"
	"log"
)

func main() {
	var flags flag.FlagSet
	procNumber, errorsBool, err := requestParser.HandleFlags(&flags)
	if err != nil {
		log.Println(err)
		return
	}
	var errors []error
	err = requestParser.ReadStdinWriteStdout(procNumber, &errors)
	if err != nil {
		log.Println(err)
	}
	if errorsBool {
		requestParser.PrintErrors(errors)
	}
}
