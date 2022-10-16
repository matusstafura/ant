package str

import (
	"regexp"
	"strings"
)

func After(s, c string) string {
	r := strings.Split(s, c)
	return string(r[1])
}

func Append(s ...string) string {
	return strings.Join(s, "")
}

func Before(s, c string) string {
	r := strings.Split(s, c)
	return string(r[0])
}

func Between(s, a, b string) string {
	re := regexp.MustCompile(`^` + a + `(.*)` + b + `$`)
	f := re.FindStringSubmatch(s)
	return f[1]
}

func EndsWith(s string, c string) bool {
	r := len(s) - len(c)
	if string(s[r:]) == c {
		return true
	}

	return false
}

func Escape(s string) string {
	m := map[string]string{"<": "&lt;", ">": "&gt;", "\"": "&quot;", "'": "&#39;"}

	re := regexp.MustCompile(`&`)
	s = re.ReplaceAllString(s, `&amp;`)

	for i, j := range m {
		re := regexp.MustCompile(i)
		s = re.ReplaceAllString(s, j)
	}

	return s
}

func Mask(s string, n int, r string) string {
	l := len(s)
	x := s[n:l]
	re := regexp.MustCompile(`.`)

	before := Limit(s, n)
	after := re.ReplaceAllString(x, r)

	return (before + after)
}

func Length(s string) int {
	return len(s)
}

func Limit(s string, n int) string {
	return s[0:n]
}

func Lower(s string) string {
	return strings.ToLower(s)
}

func LowerFirst(s string) string {
	return Lower(string(s[0])) + string(s[1:])
}

func Remove(s, c string) string {
	re := regexp.MustCompile(c)
	return re.ReplaceAllString(s, "")
}

func Replace(subject, search, replace string) string {
	re := regexp.MustCompile(search)
	return re.ReplaceAllString(subject, replace)
}

func Repeat(s string, n int) string {
	x := ""
	i := 0
	for i < n {
		x += s
		i++
	}
	return x
}

func Reverse(s string) string {
	f := make([]string, len(s))

	for i, c := range s {
		f[len(s)-1-i] = string(c)
	}
	return strings.Join(f, "")
}

func StartsWith(s string, c string) bool {
	if string(s[0:len(c)]) == c {
		return true
	}

	return false
}

func Title(s string) string {
	return strings.Title(Lower(s))
}

func Upper(s string) string {
	return strings.ToUpper(s)
}

func UpperFirst(s string) string {
	return Upper(string(s[0])) + string(s[1:])
}
