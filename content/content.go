package content

import (
	"io"
	"log"
	"net/http"
	"os"
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
type HTTPClient interface {
	Get(url string) (*http.Response, error)
}

var Client HTTPClient

func init() {
	Client = &http.Client{}
}

func FromUrl(url string) []byte {
	resp, err := Client.Get(url)
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
func SliceToFile(filename string, s []string) error {
	f, err := os.Create(filename)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	for _, v := range s {
		_, err := f.WriteString(v + SEPARATOR)
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}
