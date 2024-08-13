package server

import (
	"context"
	"encoding/json"
	"net/http"
)

type (
	IMuxContext interface {
		Context() context.Context
		JSON(statusCode int, data any)
		// Decode(data any) error
		//	GetResponseWriter() http.ResponseWriter
		//	GetRequestReader() *http.Request
		//	GetQuery(param string) string
		//	GetParam(param string) string
		//	Validate(input any) error
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

func (c *muxContext) JSON(statusCode int, payload any) {
	c.w.Header().Set("Content-Type", "applicaion/json")
	c.w.WriteHeader(statusCode)
	response, _ := json.Marshal(payload)
	c.w.Write(response)
}
