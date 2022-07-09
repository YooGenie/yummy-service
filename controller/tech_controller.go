package controller

import (
	requestDto "github.com/YooGenie/daily-work-log-service/dto/request"
	"github.com/YooGenie/daily-work-log-service/tech/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type TechController struct {
}

func (controller TechController) Init(g *echo.Group) {
	g.POST("", controller.Create)
	g.GET("/:id", controller.GetTech)
}

func (TechController) Create(ctx echo.Context) error {
	var tech requestDto.TechCreate

	if err := ctx.Bind(&tech); err != nil {
		return err
	} else {
		if err := tech.Validate(ctx); err != nil {
			return err
		}
	}

	err := service.TechService().Create(ctx, tech)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusCreated)
}


func (TechController) GetTech(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return err
	}

	member, err := service.TechService().GetTech(ctx, id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, member)

}