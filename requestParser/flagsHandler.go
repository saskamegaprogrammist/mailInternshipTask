package requestParser

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

// flags assignment

func argsAssign(flags *flag.FlagSet) {
	flags.Uint64("k", ProcNumStandart, "max number of goroutines for requests")
	flags.Bool("rootDir", false, "set root directory as files directory")
	flags.Bool("e", false, "print errors")
}

// flags lookup

func getArguments(flags *flag.FlagSet) (int, bool, bool, error) {
	var procNumber int
	var errorsBool bool
	var directoryBool bool
	numberStr := flags.Lookup("k").Value.String()
	procNumber, err := strconv.Atoi(numberStr)
	if  err != nil || procNumber == 0 {
		return procNumber, errorsBool, directoryBool, fmt.Errorf("wrong k parameter, must be positive integer")
	}
	errorsStr := flags.Lookup("e").Value.String()
	if errorsStr == "true" {
		errorsBool = true
	} else {
		errorsBool = false
	}
	dirStr := flags.Lookup("rootDir").Value.String()
	if dirStr == "true" {
		directoryBool = true
	} else {
		directoryBool = false
	}
	return procNumber, errorsBool, directoryBool, nil
}

// printing errors

func PrintErrors(errors []error) {
	for _, e := range errors {
		fmt.Println(e)
	}
}

// setting root directory

func setDir(dirBool bool) error {
	if dirBool {
		err := os.Chdir("/")
		if err != nil {
			return fmt.Errorf("error setting files directory: %s", err.Error())
		}
	}
	return nil
}

// handling flags

func HandleFlags(flags *flag.FlagSet) (int, bool, error) {
	argsAssign(flags)
	err := flags.Parse(os.Args[1:])
	if err != nil {
		return ProcNumStandart, false, err
	}
	procNumber, errorsBool, dirBool, err := getArguments(flags)
	if err != nil {
		return procNumber, errorsBool, err
	}
	err = setDir(dirBool)
	if err != nil {
		return procNumber, errorsBool, err
	}
	return procNumber, errorsBool, nil
}
