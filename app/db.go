package app

import (
	"database/sql"
	"time"

	"github.com/juandaantoniusapakpahan/go-restful-api/helper"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@root123@tcp(localhost:3306)/belajar_golang_result_api")
	helper.PanicIfError(err)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}
