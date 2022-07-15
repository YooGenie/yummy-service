package controller

import (
	requestDto "github.com/YooGenie/daily-work-log-service/dto/request"
	"github.com/YooGenie/daily-work-log-service/work/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type WorkController struct {
}

func (controller WorkController) Init(g *echo.Group) {
	g.POST("", controller.Create)

}

func (WorkController) Create(ctx echo.Context) error {
	var work requestDto.WorkCreate

	if err := ctx.Bind(&work); err != nil {
		return err
	} else {
		if err := work.Validate(ctx); err != nil {
			return err
		}
	}

	err := service.WorkService().Create(ctx, work)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusCreated)
}
