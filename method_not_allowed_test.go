package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestMethodNotAllowedHandler(t *testing.T)  {
	router := httprouter.New()

	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Method tidak diizinkan")
	})

	router.GET("/hello", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello World")
	})

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/hello", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	 response := recorder.Result()
	 body, _ := io.ReadAll(response.Body)
	 fmt.Print(string(body))
	 assert.Equal(t, "Method tidak diizinkan", string(body))

}