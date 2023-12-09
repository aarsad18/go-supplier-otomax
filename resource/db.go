package resource

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	postgres "gorm.io/driver/postgres"
	gorm "gorm.io/gorm"
)

type Notification struct {
	Channel string
	Payload string
}

type DBConn struct {
	DBSql  *sql.DB
	DBGorm *gorm.DB
	Dsn    string
}

func NewDBConn() *DBConn {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", viper.Get("DB_HOST"), viper.GetInt("DB_PORT"), viper.Get("DB_USER"), viper.Get("DB_PASS"), viper.Get("DB_NAME"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	sqlDB.SetMaxOpenConns(5)
	sqlDB.SetMaxIdleConns(1)
	sqlDB.SetConnMaxIdleTime(time.Duration(time.Duration.Milliseconds(10000)))

	return &DBConn{
		DBSql:  sqlDB,
		DBGorm: db,
		Dsn:    dsn,
	}
}
