package requestParser

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"sync"
)

func parseRequest(stringRoutine chan Request, doneChannel chan EmptyStruct, resultsChannel chan Request, wg *sync.WaitGroup, errors *[]error) {
	defer wg.Done()
	var request Request
	var err error
	request = <-stringRoutine
	resourceType := getResourceType(request.resource)
	if resourceType == URL {
		request.count, err = countURL(request.resource)
	} else {
		request.count, err = countFile(request.resource)
	}
	if err != nil {
		*errors = append(*errors, err)
	}
	resultsChannel <- request
	<-doneChannel
}

func processResults(resultsChannel chan Request, requests *[]Request, wgGlobal *sync.WaitGroup) {
	for result := range resultsChannel {
		*requests = append(*requests, result)
	}
	wgGlobal.Done()
}


func processRequests(requestChannel chan Request, resultsChannel chan Request, wgGlobal *sync.WaitGroup, procNumber int, errors *[]error) {
	doneChannel := make(chan EmptyStruct, procNumber)
	wg := &sync.WaitGroup{}
	for request := range requestChannel {
		doneChannel <- EmptyStruct{}
		stringRoutine := make(chan Request, 1)
		wg.Add(1)
		stringRoutine <- request
		go parseRequest(stringRoutine, doneChannel, resultsChannel, wg, errors)
	}
	wg.Wait()
	close(resultsChannel)
	wgGlobal.Done()
}

func writeAnswers(requests []Request, writer *bufio.Writer) error {
	var err error
	sort.SliceStable(requests, func(i, j int) bool {
		return requests[i].id < requests[j].id
	})
	total := 0
	for i := range requests{
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

func ReadStdin(procNumber int, errors *[]error) error {
	var requests []Request
	id := 0
	writer := bufio.NewWriter(os.Stdout)
	requestsChannel := make(chan Request)
	resultsChannel := make(chan Request)
	wgGlobal := &sync.WaitGroup{}
	wgGlobal.Add(2)
	go processResults(resultsChannel, &requests, wgGlobal)
	go processRequests(requestsChannel, resultsChannel, wgGlobal, procNumber, errors)

	scanner := bufio.NewScanner(os.Stdin)
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
			requestsChannel <- request
			id++
		}
	}
	close(requestsChannel)
	wgGlobal.Wait()
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