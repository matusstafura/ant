package content

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var SEPARATOR = "\n"

// returns []byte from a file
func FromFile(file string) []byte {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Println(err)
	}
	return content
}

// returns []byte from a URL
func FromUrl(url string) []byte {
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		log.Println("somethings wrong with input url", url)
	}

	html, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	return html
}

// writes to a file
func SliceToFile(filename string, s []string) {
	if s != nil {
		a := strings.Join(s, SEPARATOR)
		os.WriteFile(filename, []byte(a), 0644)
	}
}
