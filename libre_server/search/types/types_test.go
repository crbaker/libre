package types

import (
	"testing"
)

func Test_FindIdentifierWithOnlyISBN10(t *testing.T) {
	volume := dummyBookWithoutISBN13()

	id, _ := volume.Identifier()

	if id != "1234567890" {
		t.Error("1234567890", id)
	}
}

func Test_FindIdentifier(t *testing.T) {
	volume := dummyBook()

	id, _ := volume.Identifier()

	if id != "1231234567890" {
		t.Error("1231234567890", id)
	}
}

func dummyBookWithoutISBN13() Volume {

	isbn10 := Identifier{Type: "ISBN_10", Identifier: "1234567890"}

	identifiers := []Identifier{isbn10}

	volume := Volume{Title: "Test Book", IndustryIdentifiers: identifiers}
	return volume
}

func dummyBook() Volume {

	isbn10 := Identifier{Type: "ISBN_10", Identifier: "1234567890"}
	isbn13 := Identifier{Type: "ISBN_13", Identifier: "1231234567890"}

	identifiers := []Identifier{isbn10, isbn13}

	volume := Volume{Title: "Test Book", IndustryIdentifiers: identifiers}
	return volume
}
