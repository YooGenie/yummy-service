package controller

import (
	requestDto "github.com/YooGenie/daily-work-log-service/dto/request"
	"github.com/YooGenie/daily-work-log-service/member/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type MemberController struct {
}

func (controller MemberController) Init(g *echo.Group) {
	g.POST("", controller.Create)

}

func (MemberController) Create(ctx echo.Context) error {
	var member requestDto.MemberCreate

	if err := ctx.Bind(&member); err != nil {
		return err
	}
	// validate 하기

	err := service.MemberService().Create(ctx, member)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusCreated)
}
