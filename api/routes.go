// api/routes.go
package api

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	bookmarkHandler := NewBookmarkHandler() // 북마크 핸들러 생성

	// 북마크 관련 라우트 설정
	r.GET("/bookmarks", bookmarkHandler.GetBookmarks)
	r.POST("/bookmarks/add", bookmarkHandler.AddBookmark)

	// 다른 핸들러가 있다면 여기에 추가
	// r.GET("/other", otherHandler.SomeOtherFunction)
}
