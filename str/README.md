# String Helpers for Go Lang

## Usage

```go
package main

import (
	"fmt"
	"github.com/matusstafura/str"
)

func main() {
	// str.After() returns everything after first occurence
	after := str.After("Do not let the behavior of others destroy your inner peace. - Dalai Lama","- ")
	fmt.Println(after) // 'Dalai Lama'

	// str.Append(...) concatenates strings
	append := str.Append("Matus", " Stafura")
	fmt.Println(append) // 'Matus Stafura'

	// str.Before() returns everything before first occurence
	before := str.Before("Do not let the behavior of others destroy your inner peace. - Dalai Lama"," -")
	fmt.Println(before) // 'Do not let the behavior of others destroy your inner peace.'

	// str.Between() returns string between two given strings
	between := str.Between("This is a string","This", "string")
	fmt.Println(between) // ' is a '

	// str.EndsWith() checks if string ends with the given target string
	endsWith := str.EndsWith("abcdef","ef")
	fmt.Println(endsWith) // true

	// str.Escape() converts 5 characters "&", "<", ">", '"', and "'" to HTML entities.
	escape := Escape("<h1>Tom & Jerry</h1>")
	fmt.Println(escape) // '&lt;h1&gt;Tom &amp; Jerry&lt;/h1&gt;'

	// str.Length() returns lenght of the given string
	length := str.Length("Acta non verba.")
	fmt.Println(length) // '15'

	// str.Limit() truncates the string to the specified length
	limit := str.Limit("This is a string!", 6)
	fmt.Println(limit) // 'This i'

	// str.Lower() converts the string to lowercase
	lower := str.Lower("JANUARY")
	fmt.Println(lower) // 'january'

	// str.LowerFirst() converts only the first char in the string to lowercase
	lowerFirst := str.LowerFirst("JANUARY")
	fmt.Println(lowerFirst) // 'jANUARY'

	// str.Mask() converts the string to lowercase
	mask := str.Mask("4242 4242 4242 4242 4242", 4, "#")
	fmt.Println(mask) // '4242####################'

	// str.Remove(string, remove) removes the given string from the string
	remove := str.Remove("Red hot chili peppers", "e")
	fmt.Println(remove) // 'Rd hot chili ppprs

	// str.Repeat() repeats the given string n times
	repeat := str.Repeat("*", 4)
	fmt.Println(repeat) // '****'

	// str.Replace(subject, search, replace string) replaces all the occurences of the given string
	replace := str.Replace("This is a string. The string is replaced.", "string", "word")
	fmt.Println(replace) // 'This is a word. The word is replaced.'

	// str.Reverse() reverses the string
	reverse := str.Reverse("Never give up")
	fmt.Println(reverse) // 'pu evig reveN'

	// str.StartsWith() checks if string starts with the given target string
	startsWith := str.startsWith("abcdef","ab")
	fmt.Println(startsWith) // true

	// str.Title() converts string to Title case
	title := str.Title("This is a string")
	fmt.Println(title) // 'This Is A String'

	// str.Upper() converts the string to uppercase
	upper := str.Upper("Be the friend you wish you had.")
	fmt.Println(upper) // 'BE THE FRIEND YOU WISH YOU HAD.'

	// str.UpperFirst() converts only the first character in the string to uppercase
	upperFirst := str.UpperFirst("sverige")
	fmt.Println(upperFirst) // 'Sverige'
}
```

## Tests

```bash
go test ./... -v
```

## Contributions

Contributions are welcome and will be fully credited.

## License

This project is open-sourced software licensed under the MIT license.
