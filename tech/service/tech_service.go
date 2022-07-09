package service

import (
	requestDto "github.com/YooGenie/daily-work-log-service/dto/request"
	responseDto "github.com/YooGenie/daily-work-log-service/dto/response"
	"github.com/YooGenie/daily-work-log-service/tech/repository"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"sync"
)

var (
	techServiceOnce     sync.Once
	techServiceInstance *techService
)

func TechService() *techService {
	techServiceOnce.Do(func() {
		techServiceInstance = &techService{}
	})
	return techServiceInstance
}

type techService struct {
}

func (techService) Create(ctx echo.Context, creation requestDto.TechCreate) (err error) {

	if err = repository.TechRepository().Create(ctx, creation); err != nil {
		return
	}
	return

}

func (techService) GetTech(ctx echo.Context, id int64) (techSummary responseDto.TechSummary, err error) {
	techSummary, err = repository.TechRepository().GetTech(ctx, id)
	if err != nil {
		return
	}

	return
}

func (techService) GetTechs(ctx echo.Context, searchParams requestDto.SearchTechQueryParams) ([]responseDto.TechSummary, error) {
	log.Traceln("")

	return repository.TechRepository().FindAll(ctx, searchParams)
}
