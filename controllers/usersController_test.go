package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
    gin.SetMode(gin.TestMode)
    r := gin.Default()
    r.GET("/hello", HelloWorld)
    return r
}

func TestHelloWorld(t *testing.T) {
    r := setupRouter()

    req, _ := http.NewRequest("GET", "/hello", nil)
    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusOK, resp.Code)
    assert.Equal(t, "Hello, World!", resp.Body.String())
}