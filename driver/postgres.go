package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/lib/pq"
	"github.com/subosito/gotenv"
)

var db *sql.DB

func init() {
	gotenv.Load()
}
func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// ConnectDB establish a db connection
func ConnectDB() *sql.DB {
	fmt.Println(os.Getenv("ELEPHANTSQL_URL"))
	pgURL, err := pq.ParseURL("postgres://rtdwokqq:OovGoBm0nzfyOXB7Ul4JUe44eGryE6el@kandula.db.elephantsql.com:5432/rtdwokqq")
	logFatal(err)

	db, err = sql.Open("postgres", pgURL)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	return db
}
