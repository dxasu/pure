package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var SQLiteDB *sql.DB

func InitSQLite(dbPath string) error {
	var err error
	SQLiteDB, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	// 配置连接池
	SQLiteDB.SetMaxOpenConns(1) // SQLite 不支持并发写
	if err = SQLiteDB.Ping(); err != nil {
		return err
	}
	return nil
}
