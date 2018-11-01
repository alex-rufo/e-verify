package parse

import (
	"bufio"
	"net/http"
	"net/url"
	"os"
)

// From reads a file and returns the lines in a string array
func From(file string) ([]string, error) {
	_, err := url.ParseRequestURI(file)
	if err != nil {
		return loadFile(file)
	}

	return loadRemoteFile(file)

}

func loadFile(filename string) ([]string, error) {
	var lines []string
	file, err := os.Open(filename)
	if err != nil {
		return lines, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func loadRemoteFile(url string) ([]string, error) {
	var lines []string
	resp, err := http.Get(url)
	if err != nil {
		return lines, err
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
