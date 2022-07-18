package repository

import (
	common "github.com/YooGenie/daily-work-log-service/common"
	"github.com/YooGenie/daily-work-log-service/common/errors"
	requestDto "github.com/YooGenie/daily-work-log-service/dto/request"
	responseDto "github.com/YooGenie/daily-work-log-service/dto/response"
	"github.com/YooGenie/daily-work-log-service/member/entity"
	"github.com/go-xorm/xorm"
	echo "github.com/labstack/echo/v4"
	"sync"
	"time"
)

var (
	memberRepositoryOnce     sync.Once
	memberRepositoryInstance *memberRepository
)

func MemberRepository() *memberRepository {
	memberRepositoryOnce.Do(func() {
		memberRepositoryInstance = &memberRepository{}
	})
	return memberRepositoryInstance
}

type memberRepository struct {
}



func (memberRepository) Create(ctx echo.Context, creation requestDto.MemberCreate) error {

	member := entity.Member{
		Email:     creation.Email,
		Password:  common.SetEncrypt(creation.Password),
		Name:      creation.Name,
		Mobile:    creation.Mobile,
		Role:      creation.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if _, err := common.GetDB(ctx).Insert(&member); err != nil {
		return err
	}

	return nil
}

func (memberRepository) GetMember(ctx echo.Context, id int64) (memberSummary responseDto.MemberSummary, err error) {
	memberSummary.Id  = id

	queryBuilder := func() xorm.Interface {
		q := common.GetDB(ctx).Table("members")
		q.Where("1=1")
		q.And("members.id =?", id)
		return q
	}

	has, err := queryBuilder().Get(&memberSummary)
	if err != nil {
		return
	}

	if has == false {
		err = errors.NoResultError(errors.MessageNoDataFound)
		return
	}

	return
}

func (memberRepository) GetMemberByEmail(ctx echo.Context, email string) (memberSummary responseDto.MemberSummary, err error) {
	memberSummary.Email = email

	queryBuilder := func() xorm.Interface {
		q := common.GetDB(ctx).Table("members").Select("id, email ,password, name ,role")
		q.Where("1=1")
		//q.And("members.email =?", email)
		return q
	}

	has, err := queryBuilder().Get(&memberSummary)
	if err != nil {
		return
	}

	if has == false {
		err = errors.NoResultError(errors.MessageNoDataFound)
		return
	}

	return
}
