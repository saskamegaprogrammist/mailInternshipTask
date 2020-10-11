package requestParser

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// getting http request and reading response body

func countURL(url string) (int, error) {
	count := 0
	client := http.Client{
		Timeout: UrlTimeout,
	}
	response, err := client.Get(url)
	if err != nil {
		return count, fmt.Errorf("error getting url %s request: %s", url, err.Error())
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return count, fmt.Errorf("error getting url %s request, reponse status code: %d", url, response.StatusCode)
	}

	for {
		responseBuffer := make([]byte, MaxResponseBufferSize)
		_, err = response.Body.Read(responseBuffer)
		if err == io.EOF {
			count += strings.Count(string(responseBuffer), GolangString)
			break
		}
		if err != nil {
			return count, fmt.Errorf("error reading response from %s: %s", url, err.Error())
		}
		count += strings.Count(string(responseBuffer), GolangString)
	}

	return count, nil
}

// reading file

func countFile(filename string) (int, error) {
	count := 0
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return count, fmt.Errorf("error reading file %s: %s", filename, err.Error())
		}
	}
	file, err := os.Open(filename)
	if err != nil {
		return count, fmt.Errorf("error reading file %s: %s", filename, err.Error())
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return count, fmt.Errorf("error reading file %s: %s", filename, err.Error())
	}

	fileBuffer := make([]byte, fileInfo.Size())

	_, err = file.Read(fileBuffer)
	if err != nil {
		return count, fmt.Errorf("error reading file %s: %s", filename, err.Error())
	}
	count = strings.Count(string(fileBuffer), GolangString)
	return count, nil
}
