package lib

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 42069
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "postgres"
)

type Image struct {
	id    int
	name  string
	bytes []byte
}

func psqlInfo() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
}

func TestDbConnection() {
	db, err := sql.Open("postgres", psqlInfo())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

func GetImageBytesById(id int) []byte {
	db, err := sql.Open("postgres", psqlInfo())

	sqlStatement := `SELECT * FROM image WHERE id=$1;`
	row := db.QueryRow(sqlStatement, id)

	var image Image
	err = row.Scan(&image.id, &image.name, &image.bytes)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return nil
	case nil:
		return image.bytes
	default:
		panic(err)
	}
}
