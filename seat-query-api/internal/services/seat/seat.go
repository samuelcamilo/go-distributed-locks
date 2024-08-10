package seat

import (
	"context"

	"example.com/seat-query-api/pkg/logger"
)

type (
	IService interface {
		FindById(ctx context.Context, id int) (err error)
		Delete(ctx context.Context, id int) (err error)
	}

	services struct {
		log logger.Logger
	}
)

func New(log logger.Logger) IService {
	return &services{log: log}
}

func (s *services) FindById(ctx context.Context, id int) (err error) {
	return
}

func (s *services) Delete(ctx context.Context, id int) (err error) {
	return
}
