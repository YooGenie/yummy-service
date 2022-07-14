package controller

import (
	requestDto "github.com/YooGenie/daily-work-log-service/dto/request"
	"github.com/YooGenie/daily-work-log-service/tech/service"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type TechController struct {
}

func (controller TechController) Init(g *echo.Group) {
	g.POST("", controller.Create)
	g.PUT("/:id", controller.Update)
	g.GET("/:id", controller.GetTech)
	g.GET("", controller.GetTechs)
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

func (TechController) Update(ctx echo.Context) error {

	var tech requestDto.TechCreate

	techID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return err
	}

	if err := ctx.Bind(&tech); err != nil {
		return err
	} else {
		if err := tech.Validate(ctx); err != nil {
			return err
		}
	}

	tech.ID = techID

	err = service.TechService().Update(ctx, tech)
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

	tech, err := service.TechService().GetTech(ctx, id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, tech)

}

func (TechController) GetTechs(ctx echo.Context) error {
	log.Traceln("")

	searchParams := requestDto.SearchTechQueryParams{}

	if err := ctx.Bind(&searchParams); err != nil {
		return err
	}

	result, err := service.TechService().GetTechs(ctx, searchParams)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, result)
}
