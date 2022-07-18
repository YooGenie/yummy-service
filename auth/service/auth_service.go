package service

import (
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
	//비밀번호 유효성

	token, err = middleware.JwtAuthentication{}.GenerateJwtToken(middleware.JwtClaim{
		ID:   memberEntity.Id,
		Name: memberEntity.Name,
		Role: memberEntity.Role,
	})

	return
}
