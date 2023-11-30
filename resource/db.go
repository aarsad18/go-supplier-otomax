package resource

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	postgres "gorm.io/driver/postgres"
	gorm "gorm.io/gorm"
)

type Notification struct {
	Channel string
	Payload string
}

const (
	host     = "172.27.29.121"
	port     = 5432
	user     = "pl_pgsupunipin"
	password = "plus2019link"
	dbname   = "new_pos_system"
)

type DBConn struct {
	DB  *sql.DB
	Dsn string
}

func NewDBConn() *DBConn {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

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
		DB:  sqlDB,
		Dsn: dsn,
	}
}
