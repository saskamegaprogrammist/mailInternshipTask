package requestParser

import (
	"net/url"
	"time"
)

const (
	ErrorId = -1
	ProcNumStandart = 5
	GolangString    = "Go"
	FILE            = iota
	URL
	UrlTimeout = 10 * time.Second
	MaxResponseBufferSize = 128*1024
)

// resource type checker

func getResourceType(resource string) int {
	parsedURL, err := url.ParseRequestURI(resource)
	if err != nil {
		return FILE
	}
	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		return FILE
	}
	return URL
}
