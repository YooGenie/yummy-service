package controller

import (
	"github.com/YooGenie/daily-work-log-service/auth/service"
	requestDto "github.com/YooGenie/daily-work-log-service/dto/request"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthController struct {
}

func (controller AuthController) Init(g *echo.Group) {
	g.POST("/login", controller.AuthWithEmailAndPassword)
}

func (AuthController) AuthWithEmailAndPassword(ctx echo.Context) (err error) {
	var signIn requestDto.SignIn
	if err = ctx.Bind(&signIn); err != nil {
		return err
	}

	if err = signIn.Validate(ctx); err != nil {
		return err
	}

	jwtToken, err := service.AuthService().AuthWithSignIdPassword(ctx, signIn)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	//refreshToken, err := ctx.Cookie("refreshToken")
	//if err != nil || len(refreshToken.Value) == 0 {
	//	cookie := new(http.Cookie)
	//	cookie.Name = "refreshToken"
	//	cookie.Value = jwtToken.RefreshToken
	//	cookie.HttpOnly = true
	//	cookie.Path = "/"
	//	ctx.SetCookie(cookie)
	//} else {
	//	refreshToken.Value = jwtToken.RefreshToken
	//	refreshToken.HttpOnly = true
	//	refreshToken.Path = "/"
	//	ctx.SetCookie(refreshToken)
	//}

	result := map[string]string{}
	result["accessToken"] = jwtToken.AccessToken
	return ctx.JSON(http.StatusOK, result)
}
