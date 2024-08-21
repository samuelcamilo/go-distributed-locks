package server

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type (
	IMuxContext interface {
		Context() context.Context
		PathIntValue(prefix string) (value int, err error)
		Ok(data any)
		BadRequest(err error)
		// Decode(data any) error
		//	GetResponseWriter() http.ResponseWriter
		//	GetRequestReader() *http.Request
		//	GetQuery(param string) string
		//	GetParam(param string) string
		//	Validate(input any) error
	}
	ErrorResponse struct {
		Error struct {
			Code      int       `json:"code"`
			Message   string    `json:"message"`
			Details   string    `json:"details"`
			Timestamp time.Time `json:"timestamp"`
		} `json:"error"`
	}
)

type muxContext struct {
	w http.ResponseWriter
	r *http.Request
}

func newMuxContext(w http.ResponseWriter, r *http.Request) IMuxContext {
	return &muxContext{
		w: w,
		r: r,
	}
}

func (c *muxContext) Context() context.Context {
	return c.r.Context()
}

func (c *muxContext) PathIntValue(prefix string) (value int, err error) {
	value, err = strconv.Atoi(c.r.PathValue(prefix))
	return
}

func (c *muxContext) Ok(payload any) {
	c.w.Header().Set("Content-Type", "applicaion/json")
	c.w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(payload)
	c.w.Write(response)
}

func (c *muxContext) BadRequest(err error) {
	c.w.Header().Set("Content-Type", "applicaion/json")
	c.w.WriteHeader(http.StatusBadRequest)
	payload := ErrorResponse{
		Error: struct {
			Code      int       `json:"code"`
			Message   string    `json:"message"`
			Details   string    `json:"details"`
			Timestamp time.Time `json:"timestamp"`
		}{Code: http.StatusBadRequest, Message: "Bad Request", Details: err.Error(), Timestamp: time.Now()},
	}
	response, _ := json.Marshal(payload)
	c.w.Write(response)
}
