package service

import (
	requestDto "github.com/YooGenie/daily-work-log-service/dto/request"
	responseDto "github.com/YooGenie/daily-work-log-service/dto/response"
	"github.com/YooGenie/daily-work-log-service/work/repository"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
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

func (workService) GetWork(ctx echo.Context, id int64) (workSummary responseDto.WorkSummary, err error) {
	workSummary, err = repository.WorkRepository().GetWork(ctx, id)
	if err != nil {
		return
	}

	return
}

func (workService) GetWorks(ctx echo.Context, searchParams requestDto.SearchWorkQueryParams) ([]responseDto.WorkSummary, error) {
	log.Traceln("")

	return repository.WorkRepository().FindAll(ctx, searchParams)
}
