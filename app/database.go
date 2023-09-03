package app

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func Database() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/schoolshare?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/schoolshare?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
