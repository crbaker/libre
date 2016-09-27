package database

import (
	"database/sql"
	"io/ioutil"
	"os"

	pb "github.com/crbaker/libre/libre"

	// used for the sqlite3 connection
	_ "github.com/mattn/go-sqlite3"
)

const (
	sqliteDb   = "./libre_database.db"
	initScript = "./initial.sql"
)

// PersistBook saves a book to the database
func PersistBook(book *pb.Book) pb.SaveBookReply_ErrorCode {
	db, err := sql.Open("sqlite3", sqliteDb)
	checkErr(err)

	stmt, err := db.Prepare("INSERT INTO books(title) values(?)")
	checkErr(err)

	_, err = stmt.Exec(book.Title)
	checkErr(err)

	return pb.SaveBookReply_DUPLICATE
}

// InitDatabase gets a connection to the sqlite database after initializing it if needed
func InitDatabase(script ...string) {

	if !checkDatabase() {
		scriptPath := initScript
		if len(script) == 1 {
			scriptPath = script[0]
		}

		sql, err := readInitalScript(scriptPath)
		checkErr(err)

		db := openDatabase()

		_, err = db.Exec(sql)
		checkErr(err)
	}
}

func openDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", sqliteDb)
	checkErr(err)

	return db
}

func readInitalScript(script string) (string, error) {
	dat, err := ioutil.ReadFile(script)
	if err != nil {
		return "", err
	}

	return string(dat), nil
}

func checkDatabase() bool {
	if _, err := os.Stat(sqliteDb); os.IsNotExist(err) {
		return false
	}
	return true
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
