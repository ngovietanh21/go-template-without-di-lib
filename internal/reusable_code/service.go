package reusablecode

import (
	"promotion/pkg/failure"
	"promotion/pkg/logger"

	"github.com/gin-gonic/gin"
)

type Service struct {
	log  *logger.Logger
	repo *Repo
}

func newService(log *logger.Logger, repo *Repo) *Service {
	return &Service{log, repo}
}

func (s *Service) GetByCode(ctx *gin.Context, code string) (*ReusableCode, error) {
	rc, err := s.repo.GetByCode(ctx, code)
	if err != nil {
		return nil, failure.ErrorWithTrace(err)
	}
	return rc, nil
}
