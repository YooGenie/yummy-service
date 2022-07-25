package dto

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type WorkCreate struct {
	ID          int64  `json:"id" `
	Date        string `json:"date" validate:"required"`
	ProjectName string `json:"projectName" validate:"required"`
	TechID      int64  `json:"techID" validate:"required"`
	TechName    string `json:"techName" validate:"required"`
	LinkURL     string `json:"linkURL"`
	Content     string `json:"content" validate:"required"`
	Retrospect  string `json:"retrospect"`
}

func (v WorkCreate) Validate(ctx echo.Context) error {
	log.Traceln("")

	if err := ctx.Validate(v); err != nil {
		return err
	}

	return nil
}

type SearchWorkQueryParams struct {
	Date string `query:"date"`
}
