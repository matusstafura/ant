# Ant <img src="https://img.icons8.com/external-flaticons-lineal-color-flat-icons/64/000000/external-ant-animal-flaticons-lineal-color-flat-icons.png" style="height:36px"/>

Various Helpers in GoLang

```go
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
```

## TESTING

run all tests

```shell
go test ./... -v
```

coverage

```shell
./scripts.sh coverage
```
