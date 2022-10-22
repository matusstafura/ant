package extract

import (
	"testing"
)

func TestExtractUrlsFromString(t *testing.T) {
	xml := `"<?xml version="1.0" encoding="UTF-8"?>
	<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
	   <url>
		  <loc>http://www.example.com/</loc>
		  <lastmod>2005-01-01</lastmod>
		  <changefreq>monthly</changefreq>
		  <priority>0.8</priority>
	   </url>
	</urlset>"`

	got := Urls(xml)

	first := "http://www.sitemaps.org/schemas/sitemap/0.9"
	second := "http://www.example.com/"

	if got[0] != first {
		t.Fatalf("got %q, wanted %q", got, first)
	}

	if got[1] != second {
		t.Fatalf("got %q, wanted %q", got, second)
	}
}

func TestExtractUrlsFromFile(t *testing.T) {
	testFilename := "../content/test/sitemap.xml"

	got := UrlsFromFile(testFilename)
	first := "http://www.sitemaps.org/schemas/sitemap/0.9"

	if got[0] != first {
		t.Fatalf("got %q, wanted %q", got, first)
	}

	count := len(got)
	wanted := 4

	if count != wanted {
		t.Fatalf("got %v, wanted %v", count, wanted)
	}
}
