package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var PostgreSQLDB *sql.DB

func InitPostgreSQL(connStr string, maxOpenConns int) error {
	var err error
	PostgreSQLDB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("postgres open failed: %v", err)
	}
	// 配置连接池
	PostgreSQLDB.SetMaxOpenConns(maxOpenConns)
	if err = PostgreSQLDB.Ping(); err != nil {
		return fmt.Errorf("postgres ping failed: %v", err)
	}
	return nil
}
