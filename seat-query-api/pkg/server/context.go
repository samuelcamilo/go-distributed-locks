package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type (
	IMuxContext interface {
		Context() context.Context
		PathIntValue(prefix string) (int, error)
		OK(data any)
		BadRequest(err error)
		NotFound(err error)
		InternalServerError(err error)
		// Decode(data any) error
		//	GetResponseWriter() http.ResponseWriter
		//	GetRequestReader() *http.Request
		//	GetQuery(param string) string
		//	GetParam(param string) string
		//	Validate(input any) error
	}
	ErrorResponse struct {
		Code      int       `json:"code"`
		Details   string    `json:"details"`
		Timestamp time.Time `json:"timestamp"`
	}
)

type muxContext struct {
	w http.ResponseWriter
	r *http.Request
}

func newMuxContext(w http.ResponseWriter, r *http.Request) IMuxContext {
	w.Header().Set("Content-Type", "application/json")
	return &muxContext{
		w: w,
		r: r,
	}
}

func (c *muxContext) Context() context.Context {
	return c.r.Context()
}

func (c *muxContext) PathIntValue(prefix string) (int, error) {
	prefixValue := c.r.PathValue(prefix)
	value, err := strconv.Atoi(prefixValue)
	if err != nil {
		return 0, fmt.Errorf("Failed to parse '%s' from path to integer", prefixValue)
	}
	return value, err
}

func (c *muxContext) OK(data any) {
	response, _ := json.Marshal(data)

	c.w.WriteHeader(http.StatusOK)
	c.w.Write(response)
}

func (c *muxContext) BadRequest(err error) {
	statusCode := http.StatusBadRequest
	response := c.errorHandlerResponse(statusCode, err)

	c.w.WriteHeader(statusCode)
	c.w.Write(response)
}

func (c *muxContext) NotFound(err error) {
	statusCode := http.StatusNotFound
	response := c.errorHandlerResponse(statusCode, err)

	c.w.WriteHeader(statusCode)
	c.w.Write(response)
}

func (c *muxContext) InternalServerError(err error) {
	statusCode := http.StatusInternalServerError
	response := c.errorHandlerResponse(statusCode, err)

	c.w.WriteHeader(statusCode)
	c.w.Write(response)
}

func (c *muxContext) errorHandlerResponse(statusCode int, err error) []byte {
	payload := ErrorResponse{
		Code:      statusCode,
		Details:   err.Error(),
		Timestamp: time.Now(),
	}

	response, _ := json.Marshal(payload)

	return response
}
