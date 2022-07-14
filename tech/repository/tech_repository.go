package repository

import (
	common "github.com/YooGenie/daily-work-log-service/common"
	"github.com/YooGenie/daily-work-log-service/common/errors"
	requestDto "github.com/YooGenie/daily-work-log-service/dto/request"
	responseDto "github.com/YooGenie/daily-work-log-service/dto/response"
	"github.com/YooGenie/daily-work-log-service/tech/entity"
	"github.com/go-xorm/xorm"
	echo "github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
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

func (techRepository) Create(c echo.Context, creation requestDto.TechCreate) error {

	userJson, err := common.Struct2Json(common.GetUserClaim(c))
	if err != nil {
		common.Log(c).Errorln(err.Error())
		return errors.ApiInternalServerError(err.Error())
	}

	tech := entity.Tech{
		Name:    creation.Name,
		Created: []byte(userJson),
		Updated: []byte(userJson),
	}

	if _, err := common.GetDB(c).Insert(&tech); err != nil {
		return err
	}

	return nil
}

func (techRepository) Update(c echo.Context, edition requestDto.TechCreate) error {

	userJson, err := common.Struct2Json(common.GetUserClaim(c))
	if err != nil {
		common.Log(c).Errorln(err.Error())
		return errors.ApiInternalServerError(err.Error())
	}

	tech := entity.Tech{
		ID:      edition.ID,
		Name:    edition.Name,
		Updated: []byte(userJson),
	}

	if _, err := common.GetDB(c).Cols("id, name, updated").ID(edition.ID).Update(&tech); err != nil {
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

func (techRepository) FindAll(ctx echo.Context, searchParams requestDto.SearchTechQueryParams) (techSummary []responseDto.TechSummary, err error) {

	log.Traceln("")

	queryBuilder := func() xorm.Interface {
		q := common.GetDB(ctx).Table("techs")
		q.Where("1=1")

		return q
	}

	if err = queryBuilder().Desc("techs.id").Find(&techSummary); err != nil {
		return
	}

	return
}
