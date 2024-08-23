package seats

import (
	seats "example.com/seat-query-api/internal/services/seat"
	"example.com/seat-query-api/pkg/server"
)

func responseHandler(c server.IMuxContext, err error) {
	switch err {
	case seats.ErrorSessionNotFound:
		c.NotFound(err)
	default:
		c.InternalServerError(err)
	}
}
