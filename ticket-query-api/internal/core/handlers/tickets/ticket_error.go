package tickets

import (
	"example.com/ticket-query-api/internal/services/tickets"
	"example.com/ticket-query-api/pkg/server"
)

func responseHandler(c server.IMuxContext, err error) {
	switch err {
	case tickets.ErrorTicketNotFound:
		c.NotFound(err)
	default:
		c.InternalServerError(err)
	}
}
