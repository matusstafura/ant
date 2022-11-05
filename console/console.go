package console

import (
	"flag"
	"log"

	"github.com/matusstafura/ant/content"
	"github.com/matusstafura/ant/extract"
	"github.com/matusstafura/ant/str"
)

func Main() {
	var f, i, o string

	function := flag.String("f", f, "Function to run. Available: extract")
	input := flag.String("input", i, "input file or url")
	output := flag.String("output", o, "output file")

	flag.Parse()

	if *function == "" {
		log.Panic("No function specified")
	}

	if *input == "" {
		log.Panic("Input is required")
	}

	if *output == "" {
		log.Panic("Output is required")
	}

	if *function == "extract" {
		log.Println("starting extract")
		if str.StartsWith(*input, "http") {
			links := extract.UrlsFromUrl(*input)
			content.SliceToFile(*output, links)
		} else {
			links := extract.UrlsFromFile(*input)
			content.SliceToFile(*output, links)
		}
		log.Printf("finished: %s, input: %s, output: %s", *function, *input, *output)
	}
}
