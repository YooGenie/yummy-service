package dto

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type TechCreate struct {
	Name string `json:"name" validate:"required"`
}

func (v TechCreate) Validate(ctx echo.Context) error {
	log.Traceln("")

	if err := ctx.Validate(v); err != nil {
		return err
	}

	return nil
}

type SearchTechQueryParams struct {
}
