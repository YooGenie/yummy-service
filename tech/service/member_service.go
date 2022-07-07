package service

import (
	requestDto "github.com/YooGenie/daily-work-log-service/dto/request"
	"github.com/YooGenie/daily-work-log-service/tech/repository"
	"github.com/labstack/echo/v4"
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
