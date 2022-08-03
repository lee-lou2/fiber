package database

import (
	"car/conf"
	"car/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strconv"
)

func ConnectDB() {
	// Sqlite
	ConnectSqliteDB()
	// Postgres
	//ConnectPostgresqlDB()

	// 데이터베이스 마이그레이션
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Phone{})
	DB.AutoMigrate(&model.Car{})
}

func ConnectSqliteDB() {
	var err error

	// 데이터베이스 생성
	DB, err = gorm.Open(sqlite.Open(conf.Config("DATABASE_HOST_SQLITE")), &gorm.Config{})
	if err != nil {
		panic("데이터베이스 생성 실패")
	}
}

func ConnectPostgresqlDB() {
	// 데이터베이스 생성
	p := conf.Config("DATABASE_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		panic("데이터베이스 포트 설정 실패")
	}
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.Config("DATABASE_HOST"),
		port,
		conf.Config("DATABASE_USER"),
		conf.Config("DATABASE_PASSWORD"),
		conf.Config("DATABASE_NAME"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("데이터베이스 연결 실패")
	}
}
