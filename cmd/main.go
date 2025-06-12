package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/seasameoil/go-pinsight.git/api"
)

func main() {
	// Gin 라우터 생성
	r := gin.Default()

	// 정적 파일 제공
	r.Static("/static", "./static") // 정적 파일 경로를 "/static"으로 설정

	// 기본 라우트 설정 (index.html 제공)
	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	// 라우터 설정
	api.SetupRoutes(r)

	log.Println("서버가 포트 8080에서 실행 중입니다...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
