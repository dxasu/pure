package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitMySQL(dsn string, maxOpenConns, maxIdleConns int) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("mysql open failed: %v", err)
	}
	// 配置连接池
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(time.Hour)
	// 测试连接
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("mysql ping failed: %v", err)
	}
	return db, nil
}
