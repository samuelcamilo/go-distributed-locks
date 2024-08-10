package models

type (
	SeatModel struct {
		Id        int       `json:"id"`
		SessionId int       `json:"sessionId"`
		Elements  []Element `json:"elements"`
	}
	Element struct {
		Row         int    `json:"row"`
		Col         int    `json:"col"`
		Code        string `json:"code"`
		Description string `json:"description"`
		Name        string `json:"name"`
		Status      int    `json:"status"`
	}
)
