package controller

import (
	requestDto "github.com/YooGenie/daily-work-log-service/dto/request"
	"github.com/YooGenie/daily-work-log-service/member/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type MemberController struct {
}

func (controller MemberController) Init(g *echo.Group) {
	g.POST("", controller.Create)
	g.GET("/:id", controller.GetMember)

}

func (MemberController) Create(ctx echo.Context) error {
	var member requestDto.MemberCreate

	if err := ctx.Bind(&member); err != nil {
		return err
	} else {
		if err := member.Validate(ctx); err != nil {
			return err
		}
	}

	err := service.MemberService().Create(ctx, member)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusCreated)
}

func (MemberController) GetMember(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return err
	}

	member, err := service.MemberService().GetMember(ctx, id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, member)

}
