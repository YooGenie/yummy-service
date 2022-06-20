package service

import (
	requestDto "github.com/YooGenie/daily-work-log-service/dto/request"
	responseDto "github.com/YooGenie/daily-work-log-service/dto/response"
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

func (memberService) GetMember(ctx echo.Context, id int64) (memberSummary responseDto.MemberSummary, err error) {
	memberSummary, err = repository.MemberRepository().GetMember(ctx, id)
	if err != nil {
		return
	}

	return
}

func (memberService) GetMemberByEmail(ctx echo.Context, email string) (memberSummary responseDto.MemberSummary, err error) {
	memberSummary, err = repository.MemberRepository().GetMemberByEmail(ctx, email)
	if err != nil {
		return
	}

	return
}
