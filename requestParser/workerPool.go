package requestParser

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// worker pool

type WorkerPool struct {
	timeout      time.Duration
	maxProcs     int64
	currentProcs int64
	wg           sync.WaitGroup
	mutex        sync.Mutex
	requestsChan chan Request
	requests     *[]Request
	errorsBool   bool
}

// add worker to pool

func (wp *WorkerPool) addWorker() {
	atomic.AddInt64(&wp.currentProcs, 1)
	wp.wg.Add(1)
	go worker(wp)
}

// remove worker from pool

func (wp *WorkerPool) removeWorker() {
	atomic.AddInt64(&wp.currentProcs, -1)
	wp.wg.Done()
}

// add request to channel

func (wp *WorkerPool) AddRequest(request Request) {
	if atomic.LoadInt64(&wp.currentProcs) < atomic.LoadInt64(&wp.maxProcs) {
		wp.addWorker()
	}
	wp.requestsChan <- request
}

// worker function

func worker(wp *WorkerPool) {
	defer wp.removeWorker()
	for {
		select {
		case request := <-wp.requestsChan:
			parseRequest(&request, wp.errorsBool)
			wp.mutex.Lock()
			*wp.requests = append(*wp.requests, request)
			wp.mutex.Unlock()
		case <-time.After(wp.timeout):
			return
		}
	}
}

//worker task to process request

func parseRequest(request *Request, errorsBool bool) {
	var err error
	resourceType := getResourceType(request.resource)
	if resourceType == URL {
		request.count, err = countURL(request.resource)
	} else {
		request.count, err = countFile(request.resource)
	}
	if err != nil && errorsBool {
		fmt.Println(err.Error())
	}
}
