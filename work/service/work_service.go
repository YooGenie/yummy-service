package service

import (
	requestDto "github.com/YooGenie/daily-work-log-service/dto/request"
	"github.com/YooGenie/daily-work-log-service/work/repository"
	"github.com/labstack/echo/v4"
	"sync"
)

var (
	workServiceOnce     sync.Once
	workServiceInstance *workService
)

func WorkService() *workService {
	workServiceOnce.Do(func() {
		workServiceInstance = &workService{}
	})
	return workServiceInstance
}

type workService struct {
}

func (workService) Create(ctx echo.Context, creation requestDto.WorkCreate) (err error) {

	if err = repository.WorkRepository().Create(ctx, creation); err != nil {
		return
	}
	return

}
