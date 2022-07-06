package dto

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type MemberCreate struct {
	Email    string `json:"email" validate:"email,lte=50,required"`
	Password string `json:"password" validate:"gte=6,lte=50,required"`
	Name     string `json:"name" validate:"lte=50,required"`
	Mobile   string `json:"mobile" validate:"mobile,required"`
}

func (v MemberCreate) Validate(ctx echo.Context) error {
	log.Traceln("")

	if err := ctx.Validate(v); err != nil {
		return err
	}

	return nil
}


type SearchMemberQueryParams struct {
}
