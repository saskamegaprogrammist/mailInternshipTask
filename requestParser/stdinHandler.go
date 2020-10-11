package requestParser

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"sync"
	"time"
)

// writing results to stdout

func writeAnswers(requests []Request, writer *bufio.Writer) error {
	var err error
	sort.SliceStable(requests, func(i, j int) bool {
		return requests[i].id < requests[j].id
	})
	total := 0
	for i := range requests {
		_, err = writer.WriteString(fmt.Sprintf("Count for %s: %d\n", requests[i].resource, requests[i].count))
		if err != nil {
			return err
		}
		total += requests[i].count
	}
	_, err = writer.WriteString(fmt.Sprintf("Total: %d\n", total))
	if err != nil {
		return err
	}
	return nil
}

// reading from stdin in loop

func read(reader io.Reader, w io.Writer, procNumber int, errorsBool bool) error {
	var requests []Request
	id := 0
	writer := bufio.NewWriter(w)

	wp := WorkerPool{
		timeout: time.Millisecond,
		maxProcs:     int64(procNumber),
		currentProcs: 0,
		wg:           sync.WaitGroup{},
		mutex:        sync.Mutex{},
		requestsChan: make(chan Request),
		requests : &requests,
		errorsBool: errorsBool,
	}

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if scanner.Err() != nil {
			return fmt.Errorf("error scanning input:%s\n", scanner.Err().Error())
		} else {
			text := scanner.Text()
			request := Request{
				id:       id,
				resource: text,
				count:    0,
			}
			wp.AddRequest(request)
			id++
		}
	}
	wp.wg.Wait()

	err := writeAnswers(requests, writer)
	if err != nil {
		return fmt.Errorf("error writing results:%s\n", err.Error())
	}
	err = writer.Flush()
	if err != nil {
		return fmt.Errorf("error writing results:%s\n", err.Error())
	}
	return nil
}

func ReadStdinWriteStdout(procNumber int, errorsBool bool) error {
	return read(os.Stdin, os.Stdout, procNumber, errorsBool)
}
