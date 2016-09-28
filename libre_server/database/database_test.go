package database

import (
	"os"
	"testing"

	pb "github.com/crbaker/libre/libre"
)

func TestMain(m *testing.M) {
	InitDatabase("../initial.sql")
	retCode := m.Run()
	os.Remove(sqliteDb)
	os.Exit(retCode)
}

func Test_SaveBook(t *testing.T) {
	dummyBook := pb.Book{Title: "Some Book",
		Description:   "Really long read about the moon",
		PublishedDate: "2016-01-01",
		SubTitle:      "Some Sub Title"}

	code := PersistBook(&dummyBook)

	if code != pb.SaveBookReply_OK {
		t.Error(code, pb.SaveBookReply_OK)
	}

	fetchedBooks := FetchBooks()

	assertIntEquals(len(fetchedBooks), 1, t)
	assertStringEquals(fetchedBooks[0].Title, "Some Book", t)
	assertStringEquals(fetchedBooks[0].Description, "Really long read about the moon", t)
	assertStringEquals(fetchedBooks[0].PublishedDate, "2016-01-01", t)
	assertStringEquals(fetchedBooks[0].SubTitle, "Some Sub Title", t)

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
