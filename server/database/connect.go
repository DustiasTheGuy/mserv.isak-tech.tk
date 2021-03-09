package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Connect to mysql database
func Connect(schema string) (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("root:password@/%s?parseTime=true", schema))
	if err != nil {
		return nil, err
	}
	// See "Important settings" section.

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
