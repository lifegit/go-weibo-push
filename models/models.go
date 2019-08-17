package models

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"go-weibo-push/pkg/conf"
	"go-weibo-push/pkg/logging/normalLogging"
	"strings"
)

var db *gorm.DB

func init() {
	if gormDB, err := createDatabase(); err == nil {
		db = gormDB
	} else {
		normalLogging.Logger.WithError(err).Fatal("create database connection failed")
	}

	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return setting.DatabaseSetting.TablePrefix + defaultTableName
	//}
	//db.SingularTable(true)

	//db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	//db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	//db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	//enable Gorm mysql log
	if flag := conf.GetBool("enable.sqlLog"); flag {
		db.LogMode(flag)
		//f, err := os.OpenFile("mysql_gorm.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		//if err != nil {
		//	logrus.WithError(err).Fatalln("could not create mysql gorm log file")
		//}
		//logger :=  New(f,"", Ldate)
		//db.SetLogger(logger)
	}
	//db.AutoMigrate()

}

//Close clear db collection
func Close() {
	if db != nil {
		db.Close()
	}

}

func createDatabase() (*gorm.DB, error) {
	dbType := conf.GetString("db.type")
	dbAddr := fmt.Sprintf("%s:%d", conf.GetString("db.addr"), conf.GetInt("db.port"))
	dbName := conf.GetString("db.database")
	dbUser := conf.GetString("db.username")
	dbPassword := conf.GetString("db.password")
	dbCharset := conf.GetString("db.charset")
	conn := ""
	switch dbType {
	case "mysql":
		conn = fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=True&loc=Local", dbUser, dbPassword, dbAddr, dbName, dbCharset)
	case "sqlite":
		conn = dbAddr
	case "mssql":
		return nil, errors.New("TODO:suport sqlServer")
	case "postgres":
		hostPort := strings.Split(dbAddr, ":")
		if len(hostPort) == 2 {
			return nil, errors.New("db.addr must be like this host:ip")
		}
		conn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", hostPort[0], hostPort[1], dbUser, dbName, dbPassword)
	default:
		return nil, fmt.Errorf("database type %s is not supported by felix db", dbType)
	}
	return gorm.Open(dbType, conn)
}

func crudOne(m interface{}, one interface{}) (err error) {
	if db.Where(m).First(one).RecordNotFound() {
		return errors.New("resource is not found")
	}
	return nil
}

func Tt() *gorm.DB {
	return db
}
