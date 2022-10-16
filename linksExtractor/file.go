package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var SEPARATOR = "\n"

// returns []byte from a file
func getContentFromFile(file string) []byte {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Println(err)
	}
	return content
}

// returns []byte from a URL
func getContentFromUrl(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	html, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	return html
}

// writes to a file
func saveSliceToFile(filename string, s []string) {
	if s != nil {
		a := strings.Join(s, SEPARATOR)
		os.WriteFile(filename, []byte(a), 0644)
	}
}
