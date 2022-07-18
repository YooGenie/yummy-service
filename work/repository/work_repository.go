package repository

import (
	common "github.com/YooGenie/daily-work-log-service/common"
	"github.com/YooGenie/daily-work-log-service/common/errors"
	requestDto "github.com/YooGenie/daily-work-log-service/dto/request"
	responseDto "github.com/YooGenie/daily-work-log-service/dto/response"
	"github.com/YooGenie/daily-work-log-service/work/entity"
	"github.com/go-xorm/xorm"
	echo "github.com/labstack/echo/v4"
	"sync"
)

var (
	workRepositoryOnce     sync.Once
	workRepositoryInstance *workRepository
)

func WorkRepository() *workRepository {
	workRepositoryOnce.Do(func() {
		workRepositoryInstance = &workRepository{}
	})
	return workRepositoryInstance
}

type workRepository struct {
}

func (workRepository) Create(c echo.Context, creation requestDto.WorkCreate) error {

	userJson, err := common.Struct2Json(common.GetUserClaim(c))
	if err != nil {
		common.Log(c).Errorln(err.Error())
		return errors.ApiInternalServerError(err.Error())
	}

	work := entity.Work{
		Date:        creation.Date,
		ProjectName: creation.ProjectName,
		TechID:      creation.TechID,
		TechName:    creation.TechName,
		LinkURL:     creation.LinkURL,
		Content:     creation.Content,
		Retrospect:  creation.Retrospect,
		Created:     []byte(userJson),
		Updated:     []byte(userJson),
	}

	if _, err := common.GetDB(c).Insert(&work); err != nil {
		return err
	}

	return nil
}

func (workRepository) GetWork(ctx echo.Context, id int64) (workSummary responseDto.WorkSummary, err error) {
	workSummary.Id = id

	queryBuilder := func() xorm.Interface {
		q := common.GetDB(ctx).Table("works")
		q.Where("1=1")
		q.And("works.id =?", id)
		return q
	}

	has, err := queryBuilder().Get(&workSummary)
	if err != nil {
		return
	}

	if has == false {
		err = errors.NoResultError(errors.MessageNoDataFound)
		return
	}

	return
}