package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitSQLite(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	// 配置连接池
	db.SetMaxOpenConns(1) // SQLite 不支持并发写
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
