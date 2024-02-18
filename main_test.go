package main

import (
	"net/http"
	"net/http/httptest"
	"rate-limiter/internal/middleware"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type fixedRateTestCases struct {
	timeDelay  time.Duration
	requestNum int
	http200Num int
}

func TestRoutes(t *testing.T) {
	r := InitRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hello", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestFixedRate(t *testing.T) {
	f1 := []fixedRateTestCases{
		{
			timeDelay:  time.Millisecond * 200,
			requestNum: 10,
			http200Num: 2,
		}, {
			timeDelay:  time.Second,
			requestNum: 10,
			http200Num: 10,
		},
	}
	r := InitRoutes()
	r.Use(middleware.RateLimit)

	for _, val := range f1 {

		var http200Count int
		for _ = range val.requestNum {

			w := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodGet, "/hello", nil)
			if err != nil {
				t.Logf("http request failed with err: %v", err)
				t.FailNow()
			}
			r.ServeHTTP(w, req)
			if w.Result().StatusCode == 200 {
				http200Count++
			}
			time.Sleep(val.timeDelay)
		}
		assert.Equal(t, val.http200Num, http200Count)
	}

}
