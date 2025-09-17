package db

const (
	// DsnSqliteFile = "sqlite:///path/to/database.db" // 磁盘文件
	// DsnSqliteMem  = "sqlite::memory:"               // 内存数据库
	// DsnMySQL      = "mysql://root:123456@127.0.0.1:3306/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	DsnMongodb = "mongodb://root:123456@127.0.0.1:27017/admin?authSource=admin"
	// DsnPostgres   = "postgres://root:123456@127.0.0.1:5432/dbname?sslmode=disable"
	DsnRedis = "redis://root:123456@127.0.0.1:6379/0"

	// use of gorm
	DsnMySQL      = "root:123456@tcp(127.0.0.1:3306)/information_schema?charset=utf8mb4&parseTime=True&loc=Local"
	DsnPostgres   = "host=127.0.0.1 user=root password=123456 dbname=mydb port=5432 sslmode=disable"
	DsnSqliteFile = "file:./data.db?cache=shared" // 磁盘文件
	DsnSqliteMem  = "file::memory:?cache=shared"
)
