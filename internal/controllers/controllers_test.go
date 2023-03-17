package controllers_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bertcurtis/go_bert_fly/internal/controllers"
	"github.com/stretchr/testify/assert"
)

func TestSoulBrother(t *testing.T) {
	// this tells go that this test can run in Parallel
	// with other t.parallel enabled unit tests
	t.Parallel()
	body := strings.NewReader("hello")
	req := httptest.NewRequest(http.MethodPost, "/calculate", body)
	w := httptest.NewRecorder()

	controllers.SoulBrother(w, req)

	response := w.Result()
	defer response.Body.Close()

	readBody, err := io.ReadAll(response.Body)
	if err != nil {
		t.Errorf("unable to read response body")
	}

	strBody := string(readBody)
	assert.Contains(t, strBody, "Request body is not valid")
}

// func TestCalculateInvalidJSON(t *testing.T) {
// 	// this tells go that this test can run in Parallel
// 	// with other t.parallel enabled unit tests
// 	t.Parallel()

// 	body := strings.NewReader("hello")
// 	req := httptest.NewRequest(http.MethodPost, "/calculate", body)
// 	w := httptest.NewRecorder()

// 	// handle the request
// 	controllers.CalculateHandler(w, req)

// 	response := w.Result()
// 	defer response.Body.Close()

// 	readBody, err := io.ReadAll(response.Body)
// 	if err != nil {
// 		t.Errorf("unable to read response body")
// 	}

// 	strBody := string(readBody)
// 	assert.Contains(t, strBody, "Request body is not valid")
// }
