package consts

import (
	"os"

	"api/database"

	"gorm.io/gorm"
)

type ProjectConstant struct {
	gorm.Model
	Key   string
	Value string
}

func GetConstant(key string, defaultValues ...string) string {
	/*
		프로젝트 상수 불러오기
	*/
	var projectConstant ProjectConstant

	db := database.DB
	db.Select("value").Where("key = ?", key).First(&projectConstant)
	value := projectConstant.Value
	// 데이터베이스에 값이 없는 경우 환경 변수에서 조회
	if value == "" {
		value = os.Getenv(key)
		db.Create(&ProjectConstant{Key: key, Value: value})
	}
	// 데이터베이스와 환경 변수 모두에 값이 없는 경우
	if value == "" && len(defaultValues) > 0 {
		value = defaultValues[0]
		db.Create(&ProjectConstant{Key: key, Value: value})
	}
	return value
}

func SetConstant(key string, value string) {
	/*
		프로젝트 상수 저장
	*/
	db := database.DB
	db.Where("key = ?", key).FirstOrCreate(&ProjectConstant{Key: key})
	db.Model(&ProjectConstant{}).Where("key = ?", key).Update("value", value)
}
