package controllers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHomeHandler(t *testing.T) {
	hc := NewHomeController()

	req, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()

	hc.HomeHandler(w, req)

	bodybytes, _ := ioutil.ReadAll(w.Body)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "0.1", string(bodybytes))
}
