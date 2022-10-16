package main

import (
	"regexp"
)

func extract(text string) []string {
	regexUrl := `(http|ftp|https):\/\/([\w_-]+(?:(?:\.[\w_-]+)+))([\w.,@?^=%&:\/~+#-]*[\w@?^=%&\/~+#-])`
	re := regexp.MustCompile(regexUrl)
	return re.FindAllString(text, -1)
}
