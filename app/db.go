package app

import (
	"fmt"
	"github.com/lifegit/go-gulu/v2/pkg/fire"
	"gorm.io/driver/mysql"
	//"gorm.io/driver/postgres"
	//"gorm.io/driver/sqlite"
	//"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
)

// init mysql
var DB *fire.Fire

func SetUpDB() {
	if db, err := OpenConnection(); err != nil {
		log.Printf("failed to connect database, got error %v", err)
		os.Exit(1)
	} else {
		DB = fire.NewInstance(db)
		sqlDB, err := db.DB()
		if err == nil {
			err = sqlDB.Ping()
		}
		if err != nil {
			log.Printf("failed to connect database, got error %v", err)
		}

		// ConnPool db conn pool
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
	}
}

func OpenConnection() (db *gorm.DB, err error) {
	gormConf := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		// https://gorm.io/docs/gorm_config.html#NamingStrategy
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			// NameReplacer:  strings.NewReplacer("Tb", ""),
		}, // 表名不加复数s
	}
	switch Global.Db.Type {
	case "mysql":
		dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", Global.Db.Username, Global.Db.Password, Global.Db.Addr, Global.Db.Port, Global.Db.Database, Global.Db.Charset)
		db, err = gorm.Open(mysql.Open(dbDSN), gormConf)
		//case "postgres":
		//	log.Println("testing postgres...")
		//	if dbDSN == "" {
		//		dbDSN = "user=gorm password=gorm dbname=gorm host=localhost port=9920 sslmode=disable TimeZone=Asia/Shanghai"
		//	}
		//	db, err = gorm.Open(postgres.New(postgres.Config{
		//		DSN:                  dbDSN,
		//		PreferSimpleProtocol: true,
		//	}), gormConf)
		//case "sqlserver":
		//	// CREATE LOGIN gorm WITH PASSWORD = 'LoremIpsum86';
		//	// CREATE DATABASE gorm;
		//	// USE gorm;
		//	// CREATE USER gorm FROM LOGIN gorm;
		//	// sp_changedbowner 'gorm';
		//	// npm install -g sql-cli
		//	// mssql -u gorm -p LoremIpsum86 -d gorm -o 9930
		//	log.Println("testing sqlserver...")
		//	if dbDSN == "" {
		//		dbDSN = "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
		//	}
		//	db, err = gorm.Open(sqlserver.Open(dbDSN), gormConf)
		//default:
		//	log.Println("testing sqlite3...")
		//	db, err = gorm.Open(sqlite.Open(filepath.Join(os.TempDir(), "gorm.db")), gormConf)
	}

	return
}
