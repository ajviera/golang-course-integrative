package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func BenchmarkConcurrency(bt *testing.B) {
	//Parallel
	bt.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				mockGetCategory("MLA3530")
			}
		},
	)
}

func BenchmarkGetCategory(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mockGetCategory("MLA3530")
	}
}

func mockGetCategory(id string) *httptest.ResponseRecorder {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/categories/:id/prices", getCategory)
	req, err := http.NewRequest(http.MethodGet, "/categories/"+id+"/prices", nil)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	return w
}
func TestGetCategorySuccess(t *testing.T) {
	w := mockGetCategory("MLA3530")
	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}

func TestGetCategoryFail(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/categories/:id/prices", func(c *gin.Context) {
		c.String(400, "Not Found")
		fail(c, nil)
	})

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/categories/MLA351010101/prices", nil)
	router.ServeHTTP(w, r)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusBadRequest, w.Code)
	}
}

func TestParsePrices(t *testing.T) {
	body := `{ "results": [ { "price": 29749 }, { "price": 29742 } ] }`
	data := Response{}
	error := json.Unmarshal([]byte(body), &data)
	if error != nil {
		t.Error("Expected 4, got ", error)
	}
	slice := parsePrices(data)
	if len(slice) != 2 {
		t.Error("Expected 2, got ", len(slice))
	}

	if slice[0] != 29749 {
		t.Error("Expected 29749, got ", slice[0])
	}

	if slice[1] != 29742 {
		t.Error("Expected 29742, got ", slice[0])
	}
}
