package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSetupRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := SetupRouter()

	req, _ := http.NewRequest(http.MethodPost, "/api/charge", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
	}
}
