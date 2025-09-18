package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitPostgreSQL(connStr string, maxOpenConns int) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("postgres open failed: %v", err)
	}
	// 配置连接池
	db.SetMaxOpenConns(maxOpenConns)
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("postgres ping failed: %v", err)
	}
	return db, nil
}
