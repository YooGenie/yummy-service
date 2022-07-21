package service

import (
	"github.com/YooGenie/daily-work-log-service/common"
	"github.com/YooGenie/daily-work-log-service/common/errors"
	requestDto "github.com/YooGenie/daily-work-log-service/dto/request"
	memberService "github.com/YooGenie/daily-work-log-service/member/service"
	"github.com/YooGenie/daily-work-log-service/middleware"
	"github.com/labstack/echo/v4"
	"sync"
)

var (
	authServiceOnce     sync.Once
	authServiceInstance *authService
)

func AuthService() *authService {
	authServiceOnce.Do(func() {
		authServiceInstance = &authService{}
	})
	return authServiceInstance
}

type authService struct {
}

func (authService) AuthWithSignIdPassword(ctx echo.Context, signIn requestDto.SignIn) (token middleware.JwtToken, err error) {
	memberEntity, err := memberService.MemberService().GetMemberByEmail(ctx, signIn.Email)
	if err != nil {
		return
	}

	if common.SetEncrypt(signIn.Password) != memberEntity.Password {
		err = errors.NoResultError(errors.MessageInputInvalid)
		return
	}

	token, err = middleware.JwtAuthentication{}.GenerateJwtToken(middleware.JwtClaim{
		ID:   memberEntity.Id,
		Name: memberEntity.Name,
		Role: memberEntity.Role,
	})

	return
}
