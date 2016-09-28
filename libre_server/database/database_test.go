package database

import (
	"os"
	"testing"

	pb "github.com/crbaker/libre/libre"
)

func TestMain(m *testing.M) {
	retCode := m.Run()
	err := os.RemoveAll(dbLocation)
	checkErr(err)
	os.Exit(retCode)
}

func Test_SaveBook(t *testing.T) {

	authors := []string{"Penn", "Teller"}
	imageLinks := []*pb.ImageLink{&pb.ImageLink{SmallThumbnail: "http", Thumbnail: "http:..."}}
	identifiers := []*pb.Identifier{&pb.Identifier{Identifier: "ABC", Type: "isbn10"}}

	dummyBook := pb.Book{Title: "Some Book",
		Description:         "Really long read about the moon",
		PublishedDate:       "2016-01-01",
		SubTitle:            "Some Sub Title",
		Authors:             authors,
		ImageLinks:          imageLinks,
		IndustryIdentifiers: identifiers}

	code := PersistBook(&dummyBook)

	if code != pb.SaveBookReply_OK {
		t.Error(code, pb.SaveBookReply_OK)
	}

	fetchedBooks := FetchBooks()

	assertIntEquals(len(fetchedBooks), 1, t)
	assertIntEquals(len(fetchedBooks[0].Authors), 2, t)
	assertIntEquals(len(fetchedBooks[0].IndustryIdentifiers), 1, t)
	assertIntEquals(len(fetchedBooks[0].ImageLinks), 1, t)
	assertStringEquals(fetchedBooks[0].Title, "Some Book", t)
	assertStringEquals(fetchedBooks[0].Description, "Really long read about the moon", t)
	assertStringEquals(fetchedBooks[0].PublishedDate, "2016-01-01", t)
	assertStringEquals(fetchedBooks[0].SubTitle, "Some Sub Title", t)

}

func Test_ExistsInSlice(t *testing.T) {
	slice := []string{"Some value", "Test Value"}

	assertBoolEquals(existsInSlice(slice, "Test Value"), true, t)
}

func assertBoolEquals(actual bool, expected bool, t *testing.T) {
	if actual != expected {
		t.Error(actual, expected)
	}
}

func assertStringEquals(actual string, expected string, t *testing.T) {
	if actual != expected {
		t.Error(actual, expected)
	}
}
func assertIntEquals(actual int, expected int, t *testing.T) {
	if actual != expected {
		t.Error(actual, expected)
	}
}
