package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DsnMysql = "root:123456@tcp(127.0.0.1:3306)/sys?charset=utf8mb4&parseTime=True"
)

var MySQLDB *sql.DB

func InitMySQL(dsn string, maxOpenConns, maxIdleConns int) error {
	var err error
	MySQLDB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("mysql open failed: %v", err)
	}
	// 配置连接池
	MySQLDB.SetMaxOpenConns(maxOpenConns)
	MySQLDB.SetMaxIdleConns(maxIdleConns)
	MySQLDB.SetConnMaxLifetime(time.Hour)
	// 测试连接
	if err = MySQLDB.Ping(); err != nil {
		return fmt.Errorf("mysql ping failed: %v", err)
	}
	return nil
}
