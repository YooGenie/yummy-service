package service

import (
	requestDto "github.com/YooGenie/daily-work-log-service/dto/request"
	"github.com/YooGenie/daily-work-log-service/member/repository"
	"github.com/labstack/echo/v4"
	"sync"
)

var (
	memberServiceOnce     sync.Once
	memberServiceInstance *memberService
)

func MemberService() *memberService {
	memberServiceOnce.Do(func() {
		memberServiceInstance = &memberService{}
	})
	return memberServiceInstance
}

type memberService struct {
}

func (memberService) Create(ctx echo.Context, creation requestDto.MemberCreate) (err error) {

	if err = repository.MemberRepository().Create(ctx, creation); err != nil {
		return
	}
	return

}
