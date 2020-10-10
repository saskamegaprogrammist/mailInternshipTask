package requestParser

import "net/url"

const (
	GolangString = "Go"
	FILE         = iota
	URL
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
