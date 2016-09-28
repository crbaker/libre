package database

import (
	"database/sql"
	"io/ioutil"
	"os"
	"strings"

	"log"

	pb "github.com/crbaker/libre/libre"

	// used for the sqlite3 connection
	_ "github.com/mattn/go-sqlite3"
)

const (
	sqliteDb   = "./libre_database.db"
	initScript = "./initial.sql"
)

var (
	currentDb   *sql.DB
	dbConnected bool
)

// PersistBook saves a book to the database
func PersistBook(book *pb.Book) pb.SaveBookReply_ErrorCode {
	db := openDatabase()

	stmt, err := db.Prepare("INSERT INTO books(title, sub_title, description, published_date) values(?,?,?,?)")
	checkErr(err)

	res, err := stmt.Exec(book.Title, book.SubTitle, book.Description, book.PublishedDate)
	checkErr(err)

	_, err = res.LastInsertId()
	checkErr(err)

	return pb.SaveBookReply_OK
}

// FetchBooks fetches the collection of books from the database
func FetchBooks() []*pb.Book {
	db := openDatabase()

	stmt, err := db.Prepare("SELECT id, title, sub_title, description, published_date FROM books")
	checkErr(err)

	rows, err := stmt.Query()

	defer rows.Close()

	var books []*pb.Book

	for rows.Next() {
		book := rowToBook(rows)
		books = append(books, &book)
	}

	return books
}

func rowToBook(rows *sql.Rows) pb.Book {
	var (
		id            int
		title         string
		subTitle      string
		description   string
		publishedDate string
	)
	rows.Scan(&id, &title, &subTitle, &description, &publishedDate)
	return pb.Book{Title: title, Description: description, PublishedDate: publishedDate, SubTitle: subTitle}
}

type byName []os.FileInfo

func (f byName) Len() int {
	return len(f)
}
func (f byName) Less(i, j int) bool {
	return strings.Compare(f[i].Name(), f[j].Name()) == -1
}
func (f byName) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func MigrateDatabase(migrationsPath ...string) {

	ioutil.readd
	files, err := ioutil.ReadDir(migrationsPath[0])
	checkErr(err)

	// for

	// db := openDatabase()

	// stmt, err := db.Prepare("SELECT id, title, sub_title, description, published_date FROM books")
	// checkErr(err)

	// rows, err := stmt.Query()
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

	if !dbConnected {
		log.Println("open database")
		var err error
		currentDb, err = sql.Open("sqlite3", sqliteDb)
		checkErr(err)

		dbConnected = true
	}

	return currentDb
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
