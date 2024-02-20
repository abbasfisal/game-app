package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// ...

type MYSQLDB struct {
	db *sql.DB
}

func New() *MYSQLDB {
	db, err := sql.Open("mysql", "root:password@(localhost:3307)/gameapp_db")
	if err != nil {
		panic(fmt.Errorf("unable to open mysql : %v", err))
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MYSQLDB{db: db}
}
