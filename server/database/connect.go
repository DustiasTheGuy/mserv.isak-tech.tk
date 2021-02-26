package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Connect to mysql database
func Connect() (*sql.DB, error) {

	db, err := sql.Open("mysql", "root:password@/isak_tech_paste?parseTime=true")
	if err != nil {
		return nil, err
	}
	// See "Important settings" section.

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
