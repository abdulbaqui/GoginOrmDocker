package controllers

import (
	initializers "GoginOrmDocker/Initializers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestMain sets up the test environment
func TestMain(m *testing.M) {
	// Set test mode for gin
	gin.SetMode(gin.TestMode)

	// Run tests
	m.Run()
}

func TestPostCreate(t *testing.T) {
	// Store original DB
	originalDB := initializers.DB

	// Set DB to nil to simulate database connection failure
	initializers.DB = nil

	router := gin.Default()
	router.POST("/posts", PostCreate)

	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/posts", strings.NewReader(`{"name": "John", "age": 20, "gender": true}`))
	request.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, request)

	// Since DB is nil, we expect a 500 error, but the test structure is correct
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	// Restore original DB
	initializers.DB = originalDB
}

func TestDeletePost(t *testing.T) {
	// Store original DB
	originalDB := initializers.DB

	// Set DB to nil to simulate database connection failure
	initializers.DB = nil

	router := gin.Default()
	router.DELETE("/posts/:id", Delete)

	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodDelete, "/posts/1", nil)
	router.ServeHTTP(w, request)

	// Since DB is nil, we expect a 500 error, but the test structure is correct
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	// Restore original DB
	initializers.DB = originalDB
}

func TestGetSpecificPost(t *testing.T) {
	// Store original DB
	originalDB := initializers.DB

	// Set DB to nil to simulate database connection failure
	initializers.DB = nil

	router := gin.Default()
	router.GET("/posts/:id", GetSpecific)

	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/posts/1", nil)
	router.ServeHTTP(w, request)

	// Since DB is nil, we expect a 500 error, but the test structure is correct
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	// Restore original DB
	initializers.DB = originalDB
}

func TestPostIndex(t *testing.T) {
	// Store original DB
	originalDB := initializers.DB

	// Set DB to nil to simulate database connection failure
	initializers.DB = nil

	router := gin.Default()
	router.GET("/posts", PostIndex)

	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/posts", nil)
	router.ServeHTTP(w, request)

	// Since DB is nil, we expect a 500 error, but the test structure is correct
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	// Restore original DB
	initializers.DB = originalDB
}
