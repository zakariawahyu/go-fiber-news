package utils

import (
	"github.com/zakariawahyu/go-fiber-news/config"
	"github.com/zakariawahyu/go-fiber-news/exception"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type Conn struct {
	Mysql *gorm.DB
}

func NewDbConnection(cfg *config.Config) *Conn {
	return &Conn{
		Mysql: InitMysql(cfg),
	}
}

func InitMysql(cfg *config.Config) *gorm.DB {
	db, err := gorm.Open(mysql.Open(cfg.DB.DSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	exception.PanicIfNeeded(err)

	sqlDB, errSqlDB := db.DB()
	exception.PanicIfNeeded(errSqlDB)
	sqlDB.SetMaxIdleConns(cfg.DB.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.DB.MaxOpenConns)
	sqlDB.SetConnMaxIdleTime(cfg.DB.ConnMaxIdleTime * time.Minute)
	sqlDB.SetConnMaxLifetime(cfg.DB.ConnMaxLifetime * time.Minute)

	errPing := sqlDB.Ping()
	exception.PanicIfNeeded(errPing)

	return db
}
