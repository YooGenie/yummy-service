package repository

import (
	common "github.com/YooGenie/daily-work-log-service/common"
	requestDto "github.com/YooGenie/daily-work-log-service/dto/request"
	"github.com/YooGenie/daily-work-log-service/member/entity"
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
		Email:    creation.Email,
		Password: creation.Password,
		Name:     creation.Name,
		Mobile:   creation.Mobile,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if _, err := common.GetDB(ctx).Insert(&member); err != nil {
		return err
	}


	return nil
}
