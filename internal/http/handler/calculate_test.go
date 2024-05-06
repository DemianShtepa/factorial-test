package handler_test

import (
	"bytes"
	"github.com/DemianShtepa/factorial-test/internal/http/handler"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalculateHandler_FailedForInvalidInput(t *testing.T) {
	calculateHandler := handler.CalculateHandler()

	t.Run("fail with empty body", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "/calculate", nil)
		response := httptest.NewRecorder()
		calculateHandler(response, request, nil)

		assert.Equal(t, http.StatusBadRequest, response.Result().StatusCode)
	})

	t.Run("fail with negative params", func(t *testing.T) {
		var buffer bytes.Buffer
		buffer.WriteString("{\"a\": -1, \"b\": 10}")
		request := httptest.NewRequest(http.MethodPost, "/calculate", &buffer)
		response := httptest.NewRecorder()
		calculateHandler(response, request, nil)

		assert.Equal(t, http.StatusBadRequest, response.Result().StatusCode)
	})
}

func TestCalculateHandler(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString("{\"a\": 5, \"b\": 10}")
	request := httptest.NewRequest(http.MethodPost, "/calculate", &buffer)
	response := httptest.NewRecorder()
	calculateHandler := handler.CalculateHandler()

	calculateHandler(response, request, nil)

	assert.Equal(t, http.StatusOK, response.Result().StatusCode)
}
