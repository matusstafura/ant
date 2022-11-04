# API

```go
	sitemap := extract.UrlsFromUrl("https://b2b.pgn.com.pl/sitemap.xml")
	content.SliceToFile("here.txt", sitemap)
```

```shell
    app urls2file -i "http://www.example.com/sitemap.xml" -o "links.txt"
```
