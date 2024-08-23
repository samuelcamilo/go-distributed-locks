package events

import (
	"example.com/event-query-api/internal/services/events"
	"example.com/event-query-api/pkg/server"
)

func responseHandler(c server.IMuxContext, err error) {
	switch err {
	case events.ErrorEventNotFound:
		c.NotFound(err)
	default:
		c.InternalServerError(err)
	}
}
