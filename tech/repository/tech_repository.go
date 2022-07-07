package repository

import (
	common "github.com/YooGenie/daily-work-log-service/common"
	requestDto "github.com/YooGenie/daily-work-log-service/dto/request"
	"github.com/YooGenie/daily-work-log-service/tech/entity"
	echo "github.com/labstack/echo/v4"
	"sync"
)

var (
	techRepositoryOnce     sync.Once
	techRepositoryInstance *techRepository
)

func TechRepository() *techRepository {
	techRepositoryOnce.Do(func() {
		techRepositoryInstance = &techRepository{}
	})
	return techRepositoryInstance
}

type techRepository struct {
}


func (techRepository) Create(ctx echo.Context, creation requestDto.TechCreate) error {

	tech := entity.Tech{
		Name:      creation.Name,
		//세션에 있는 값을 넣어야함
		//Created:
		//Updated:
	}

	if _, err := common.GetDB(ctx).Insert(&tech); err != nil {
		return err
	}

	return nil
}