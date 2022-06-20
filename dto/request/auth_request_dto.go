package dto

import "github.com/labstack/echo/v4"

type SignIn struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"gte=6,lte=100"`
}

func (s SignIn) Validate(ctx echo.Context) (err error) {
	if err = ctx.Validate(s); err != nil {
		return
	}
	return
}
