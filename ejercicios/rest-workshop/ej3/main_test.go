package main__test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"rest-workshop/ej2/src/router"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestItem struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func TestCreationHandler(t *testing.T) {
	r := router.CreateRouter()

	if r == nil {
		t.Error("Expected router to be created")
	}

	data := TestItem{
		ID:    "1",
		Name:  "testItem",
		Price: 10.0,
	}
	marshalledData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(marshalledData))
	req.Header.Set("Content-Type", "application/json")

	secondReq, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(marshalledData))
	secondReq.Header.Set("Content-Type", "application/json")

	resRecorder := httptest.NewRecorder()
	secondResRecorder := httptest.NewRecorder()

	r.ServeHTTP(resRecorder, req)
	r.ServeHTTP(secondResRecorder, secondReq)

	var result TestItem

	err := json.Unmarshal(resRecorder.Body.Bytes(), &result)

	if err != nil {
		t.Error("Expected a valid response")
	}

	assert.Equal(t, http.StatusCreated, resRecorder.Code)
	assert.Equal(t, http.StatusConflict, secondResRecorder.Code)
	assert.Equal(t, data, result)
}
