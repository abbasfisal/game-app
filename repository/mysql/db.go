package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Config struct {
	Username string
	Password string
	Port     int
	Host     string
	DbName   string
}

// ...

type MYSQLDB struct {
	config Config
	db     *sql.DB
}

func New(cfg Config) *MYSQLDB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DbName))
	if err != nil {
		panic(fmt.Errorf("unable to open mysql : %v", err))
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MYSQLDB{config: cfg, db: db}
}
