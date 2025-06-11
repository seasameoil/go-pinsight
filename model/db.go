package model

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open(sqlite.Open("pinsight.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("DB 연결 실패:", err)
	}

	// 테이블 자동 생성
	db.AutoMigrate(&Bookmark{})

	DB = db
}
