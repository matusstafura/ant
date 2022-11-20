package extract

import (
	"testing"
)

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func assertInt(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

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

	assertString(t, got[0], first)
	assertString(t, got[1], second)
}

func TestExtractUrlsFromFile(t *testing.T) {
	testFilename := "../content/test/sitemap.xml"

	got := UrlsFromFile(testFilename)
	first := "http://www.sitemaps.org/schemas/sitemap/0.9"

	assertString(t, got[0], first)

	count := len(got)
	wanted := 4

	assertInt(t, count, wanted)
}
