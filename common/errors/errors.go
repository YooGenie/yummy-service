package errors

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// 인증실패
func StatusUnauthorized(message string) error {
	return echo.NewHTTPError(http.StatusUnauthorized, message)
}

// 403 접근금지
func StatusForbidden(message string) error {
	return echo.NewHTTPError(http.StatusForbidden, message)
}

// 413 제한한 크기보다 큰 경우
func ApiRequestTooBigError(message string) error {
	return echo.NewHTTPError(http.StatusRequestEntityTooLarge, message)
}

//500
func ApiInternalServerError(message string) error {
	return echo.NewHTTPError(http.StatusInternalServerError, message)
}

// 406 : 헤더에 적혀 있는 형식을 생성해낼 수 없을 때 발생하는 에러
func ApiNotAcceptableError(message string) error {
	return echo.NewHTTPError(http.StatusNotAcceptable, message)
}

//400
func ParamsValidationError(message string) error {
	return echo.NewHTTPError(http.StatusBadRequest, message)
}

//404
func NoResultError(message string) error {
	return echo.NewHTTPError(http.StatusNotFound, message)
}
