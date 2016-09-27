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
	dummyBook := pb.Book{Title: "Some Book", Description: "Really long read about the moon"}

	PersistBook(&dummyBook)
}
