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
