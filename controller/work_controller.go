package controller

import (
	requestDto "github.com/YooGenie/daily-work-log-service/dto/request"
	"github.com/YooGenie/daily-work-log-service/work/service"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type WorkController struct {
}

func (controller WorkController) Init(g *echo.Group) {
	g.POST("", controller.Create)
	g.GET("/:id", controller.GetWork)
	g.GET("", controller.GetWorks)
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

func (WorkController) GetWork(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return err
	}

	work, err := service.WorkService().GetWork(ctx, id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, work)

}

func (WorkController) GetWorks(ctx echo.Context) error {
	log.Traceln("")

	searchParams := requestDto.SearchWorkQueryParams{}

	if err := ctx.Bind(&searchParams); err != nil {
		return err
	}

	result, err := service.WorkService().GetWorks(ctx, searchParams)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, result)
}
