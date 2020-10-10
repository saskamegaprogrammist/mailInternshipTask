package requestParser

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// getting http request and reading response body

func countURL(url string) (int, error) {
	count := 0
	response, err := http.Get(url)
	if err != nil {
		return count, fmt.Errorf("error getting url %s request: %s", url, err.Error())
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return count, fmt.Errorf("error getting url %s request, reponse status code: %d", url, response.StatusCode)
	}
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return count, fmt.Errorf("error reading response body from url %s request: %s", url, err.Error())
	}

	count = strings.Count(string(bodyBytes), GolangString)
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
