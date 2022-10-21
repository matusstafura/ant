package extract

import (
	"regexp"
)

// extracts http(s) from a given string
func Urls(text string) []string {
	regexUrl := `(http|ftp|https):\/\/([\w_-]+(?:(?:\.[\w_-]+)+))([\w.,@?^=%&:\/~+#-]*[\w@?^=%&\/~+#-])`
	re := regexp.MustCompile(regexUrl)
	return re.FindAllString(text, -1)
}
