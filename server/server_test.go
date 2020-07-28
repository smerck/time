package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// adding intial tests
func TestHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)

	rec := httptest.NewRecorder()
	handler := &timeHandler{}

	handler.ServeHTTP(rec, req)

	assert.Equal(t, rec.Code, 200)
	b, err := ioutil.ReadAll(rec.Body)
	assert.Nil(t, err)
	resp := timeResponse{}
	err = json.Unmarshal(b, &resp)
	assert.NotNil(t, resp.Time)
}
