package repository

import (
	common "github.com/YooGenie/daily-work-log-service/common"
	"github.com/YooGenie/daily-work-log-service/common/errors"
	requestDto "github.com/YooGenie/daily-work-log-service/dto/request"
	responseDto "github.com/YooGenie/daily-work-log-service/dto/response"
	"github.com/YooGenie/daily-work-log-service/tech/entity"
	"github.com/go-xorm/xorm"
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


func (techRepository) GetTech(ctx echo.Context, id int64) (techSummary responseDto.TechSummary, err error) {
	techSummary.Id = id

	queryBuilder := func() xorm.Interface {
		q := common.GetDB(ctx).Table("techs")
		q.Where("1=1")
		q.And("techs.id =?", id)
		return q
	}

	has, err := queryBuilder().Get(&techSummary)
	if err != nil {
		return
	}

	if has == false {
		err = errors.NoResultError(errors.MessageNoDataFound)
		return
	}

	return
}