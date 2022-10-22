package extract

import (
	"regexp"

	"github.com/matusstafura/ant/content"
)

// extracts http(s) from a given string
func Urls(text string) []string {
	regexUrl := `(http|ftp|https):\/\/([\w_-]+(?:(?:\.[\w_-]+)+))([\w.,@?^=%&:\/~+#-]*[\w@?^=%&\/~+#-])`
	re := regexp.MustCompile(regexUrl)
	return re.FindAllString(text, -1)
}

// extracts http(s) from a file
func UrlsFromFile(filename string) []string {
	file := content.FromFile(filename)
	return Urls(string(file))
}
