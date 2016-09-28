package database

import (
	"encoding/json"

	pb "github.com/crbaker/libre/libre"

	"github.com/HouzuoGuo/tiedot/db"
	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
)

const (
	dbLocation = "./libre_database.db"
)

// PersistBook saves a book to the database
func PersistBook(book *pb.Book) pb.SaveBookReply_ErrorCode {
	db := openDatabase()
	defer db.Close()

	coll := getCollection(db, "library")

	_, err := coll.Insert(bookToRaw(book))

	if err != nil {
		panic(err)
	}

	return pb.SaveBookReply_OK
}

// FetchBooks fetches the collection of books from the database
func FetchBooks() []*pb.Book {
	db := openDatabase()
	defer db.Close()

	coll := getCollection(db, "library")

	var books []*pb.Book

	coll.ForEachDoc(func(id int, content []byte) (willMoveOn bool) {
		var dat map[string]interface{}
		if err := json.Unmarshal(content, &dat); err != nil {
			panic(err)
		}

		book := rawToBook(dat)
		books = append(books, &book)

		return true
	})

	return books
}

func existsInSlice(slice []string, val string) bool {
	for _, b := range slice {
		if b == val {
			return true
		}
	}
	return false
}

func getCollection(db *db.DB, col string) *db.Col {
	if !existsInSlice(db.AllCols(), col) {
		db.Create(col)
	}
	return db.Use(col)
}

func bookToRaw(book *pb.Book) map[string]interface{} {
	return structs.Map(book)
}

func rawToBook(record map[string]interface{}) pb.Book {
	var book pb.Book

	err := mapstructure.Decode(record, &book)
	checkErr(err)

	return book
}

func openDatabase() *db.DB {
	db, err := db.OpenDB(dbLocation)
	if err != nil {
		panic(err)
	}

	return db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
