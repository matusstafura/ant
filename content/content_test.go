package content

import (
	"testing"
)

func TestFromFile(t *testing.T) {
	got := FromFile("test/plaintext.md")

	text := "some text\n"
	if string(got) != text {
		t.Fatalf("got %q, wanted %q", got, text)
	}
}
