package requestParser

import (
	"flag"
	"fmt"
	"os"
)


// setting root directory

func setDir(dir string) error {
	err := os.Chdir(dir)
	if err != nil {
		return fmt.Errorf("error setting files directory: %s", err.Error())
	}
	return nil
}


// flags assignment

func FlagsParse() (int, bool, error) {
	procNum := flag.Uint("k", ProcNumStandart, "max number of goroutines for requests")
	filesDirectory := flag.String("dir", ".", "set files directory")
	errorsBool := flag.Bool("e", false, "print errors")
	flag.Parse()
	if *procNum == 0 {
		return int(*procNum), *errorsBool, fmt.Errorf("wrong k parameter, must be positive integer")
	}
	err := setDir(*filesDirectory)
	if err != nil 	{
		return int(*procNum), *errorsBool, err
	}
	return int(*procNum), *errorsBool, nil
}