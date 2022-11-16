package content

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestFromFile(t *testing.T) {
	got := FromFile("test/plaintext.md")

	text := "some text\n"
	if string(got) != text {
		t.Fatalf("got %q, wanted %q", got, text)
	}
}

type mockClient struct {
}

func (m mockClient) Get(url string) (*http.Response, error) {
	return &http.Response{
		Body: ioutil.NopCloser(strings.NewReader("text")),
	}, nil
}

func TestFromUrl(t *testing.T) {
	Client = mockClient{}
	got := FromUrl("somewebsite.tld")
	want := "text"

	if string(got) != "text" {
		t.Fatalf("got %q, wanted %q", got, want)
	}
}

func TestSliceToFile(t *testing.T) {
	s := []string{"a", "b", "c"}
	SliceToFile("test/slice.txt", s)
	got := FromFile("test/slice.txt")

	text := "a\nb\nc\n"
	if string(got) != text {
		t.Fatalf("got %q, wanted %q", got, text)
	}
}

func TestSliceWrongFile(t *testing.T) {
	s := []string{"a", "b", "c"}
	err := SliceToFile("", s)
	if err != nil {
		t.Fatal(err)
	}
}
