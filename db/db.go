package db

import (
	"database/sql"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/tursodatabase/go-libsql"
)

const DefaultDBPath = "file:./db.db"

var DB *sql.DB

func New() (*sql.DB, error) {
	path, ok := os.LookupEnv("DB_PATH")
	if !ok {
		path = DefaultDBPath
	}
	println(path)
	return sql.Open("libsql", path)
}

func init() {
	var err error
	DB, err = New()
	if err != nil {
		panic(err)
	}
}
