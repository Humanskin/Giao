package sqlconfig

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

var db *sql.DB

func InitDB() (err error) {
	// Data Source Name
	dsn := "root:root@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	db, err = sql.Open("mysql",dsn)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := InitDB()
	if err != nil {
		fmt.Printf("error from %v", err)
	}
}