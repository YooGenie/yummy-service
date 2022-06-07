package service

import (
	"context"
	requestDto "github.com/YooGenie/daily-work-log-service/dto/request"
	"github.com/YooGenie/daily-work-log-service/member/repository"
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

func (memberService) Create(ctx context.Context, creation requestDto.MemberCreate) (err error) {

	if err = repository.MemberRepository().Create(ctx, creation); err != nil {
		return
	}
	return

}
