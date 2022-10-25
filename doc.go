/*
# Various Helpers in GoLang

## Usage

import (
	"fmt"
	"github.com/matusstafura/ant/extract"
)

func main() {
	sitemap := extract.UrlsFromFile("content/test/sitemap.xml")
	fmt.Println(sitemap) // [http://www.example.com http://www.other.com]

	sitemap := extract.UrlsFromUrl("https://www.example.com")
	fmt.Println(sitemap) // [http://www.example.com http://www.other.com]
}

## Tests

```shell
go test ./... -v
```

## Contributions

Contributions are welcome. Please see [CONTRIBUTING](https://github.com/matusstafura/.github/blob/main/CONTRIBUTING.md) for details.

## License

This project is open-sourced software licensed under the MIT license.
*/

package main
