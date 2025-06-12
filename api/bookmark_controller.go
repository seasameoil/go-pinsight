package api

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/seasameoil/go-pinsight.git/model"
)

type BookmarkHandler struct {
	mu        sync.Mutex
	bookmarks []model.Bookmark
}

// NewBookmarkHandler 생성자 함수
func NewBookmarkHandler() *BookmarkHandler {
	return &BookmarkHandler{
		bookmarks: []model.Bookmark{},
	}
}

// AddBookmark 핸들러: 북마크 추가
func (h *BookmarkHandler) AddBookmark(c *gin.Context) {
	var bookmark model.Bookmark

	// 요청 본문에서 JSON 데이터를 바인딩
	if err := c.ShouldBindJSON(&bookmark); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.mu.Lock()
	bookmark.ID = len(h.bookmarks) + 1
	bookmark.CreatedAt = time.Now()
	h.bookmarks = append(h.bookmarks, bookmark)
	h.mu.Unlock()

	c.JSON(http.StatusCreated, bookmark)
}

// GetBookmarks 핸들러: 모든 북마크 조회
func (h *BookmarkHandler) GetBookmarks(c *gin.Context) {
	h.mu.Lock()
	defer h.mu.Unlock()

	c.JSON(http.StatusOK, h.bookmarks)
}
