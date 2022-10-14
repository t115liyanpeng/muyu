package muyudb

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var muyudb *sqlx.DB

func InitedMuYuDB(dsn string) (err error) {
	muyudb, err = sqlx.Open("mysql", dsn)
	if err == nil {
		muyudb.SetMaxOpenConns(300)
		muyudb.SetConnMaxLifetime(10)
	}
	return err
}

func Query(query string, args ...any) (*sql.Rows, error) {
	return muyudb.Query(query, args)
}

func Select(dst interface{}, query string, args ...interface{}) error {
	return muyudb.Select(dst, query, args)
}

func Execute(query string, args ...any) (sql.Result, error) {
	return muyudb.Exec(query, args...)
}
