package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

//! unittest for HTTP connection
func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/create", ReturnRandomNumber).Methods("GET")
	return router
}

func TestReturnRandomNumber(t *testing.T) {
	request, _ := http.NewRequest("GET", "/create", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}
