package main

import (
	"flag"
	"fmt"
	"github.com/saskamegaprogrammist/mailInternshipTask/requestParser"
	"log"
	"os"
	"strconv"
)

// flags assignment

func argsAssign(flags *flag.FlagSet) {
	flags.Uint64("k", 5, "max number of goroutines for requests")
	flags.Bool("e", false, "print errors")
}

// flags lookup

func getArguments(flags *flag.FlagSet) (int, bool) {
	var errorsBool bool
	var procNumber int
	numberStr := flags.Lookup("k").Value.String()
	procNumber, err := strconv.Atoi(numberStr)
	if err != nil {
		procNumber = 5
	}
	errorsStr := flags.Lookup("e").Value.String()
	if errorsStr == "true" {
		errorsBool = true
	} else {
		errorsBool = false
	}
	return procNumber, errorsBool
}

// printing errors

func printErrors(errors []error) {
	for _, e := range errors {
		fmt.Println(e)
	}
}

func main() {
	var flags flag.FlagSet
	argsAssign(&flags)
	err := flags.Parse(os.Args[1:])
	if err != nil {
		log.Println(fmt.Errorf("flag input error").Error())
	}
	procNumber, errorsBool := getArguments(&flags)
	var errors []error
	err = requestParser.ReadStdinWriteStdout(procNumber, &errors)
	if err != nil {
		log.Println(err)
	}
	if errorsBool {
		printErrors(errors)
	}
}
