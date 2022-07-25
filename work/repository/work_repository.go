package repository

import (
	common "github.com/YooGenie/daily-work-log-service/common"
	"github.com/YooGenie/daily-work-log-service/common/errors"
	requestDto "github.com/YooGenie/daily-work-log-service/dto/request"
	responseDto "github.com/YooGenie/daily-work-log-service/dto/response"
	"github.com/YooGenie/daily-work-log-service/work/entity"
	"github.com/go-xorm/xorm"
	echo "github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
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

func (workRepository) Update(c echo.Context, edition requestDto.WorkCreate) error {

	userJson, err := common.Struct2Json(common.GetUserClaim(c))
	if err != nil {
		common.Log(c).Errorln(err.Error())
		return errors.ApiInternalServerError(err.Error())
	}

	work := entity.Work{
		Date:        edition.Date,
		ProjectName: edition.ProjectName,
		TechID:      edition.TechID,
		TechName:    edition.TechName,
		LinkURL:     edition.LinkURL,
		Content:     edition.Content,
		Retrospect:  edition.Retrospect,
		Updated:     []byte(userJson),
	}

	if _, err := common.GetDB(c).Cols("id, date,project_name, tech_id, tech_name, link_url,content,  retrospect, updated").ID(edition.ID).Update(&work); err != nil {
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

func (workRepository) FindAll(ctx echo.Context, searchParams requestDto.SearchWorkQueryParams) (workSummary []responseDto.WorkSummary, err error) {

	log.Traceln("")

	queryBuilder := func() xorm.Interface {
		q := common.GetDB(ctx).Table("works")
		q.Where("1=1")
		if searchParams.Date != "" {
			q.And("date = ?", searchParams.Date)
		}

		return q
	}

	if err = queryBuilder().Desc("works.id").Find(&workSummary); err != nil {
		return
	}

	return
}
