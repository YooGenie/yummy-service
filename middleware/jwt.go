package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/YooGenie/daily-work-log-service/common"
	error2 "github.com/YooGenie/daily-work-log-service/common/errors"
	"github.com/YooGenie/daily-work-log-service/config"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strings"
	"time"
)

var InvalidAccessToken = errors.New("invalid access token")
var AccessTokenExpired = errors.New("access token expired")

type JwtAuthentication struct {
}

//토큰 발급할 내용을 구조체 한다
type JwtClaim struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

type JwtToken struct {
	AccessToken  string
	RefreshToken string
}

func (JwtAuthentication) GenerateJwtToken(claim JwtClaim) (JwtToken, error) {
	claimMap, err := claim.ConvertMap()
	if err != nil {
		return JwtToken{}, err
	}

	accessTokenClaims := jwt.MapClaims{}
	for key, value := range claimMap {
		accessTokenClaims[key] = value
	}

	//토큰 제한시간 설정
	accessTokenClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	//토큰 만들기
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims).SignedString([]byte(config.Config.Jwt.JwtSecret))

	if err != nil {
		return JwtToken{}, err
	}

	//refreshToken 사용 보류
	//refreshTokenClaims := jwt.MapClaims{}
	//for key, value := range claimMap {
	//	refreshTokenClaims[key] = value
	//}
	//
	//refreshTokenClaims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	//refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims).SignedString([]byte(config.Config.Jwt.JwtSecret))

	return JwtToken{
		AccessToken: accessToken,
	}, nil
}

func (c JwtClaim) ConvertMap() (map[string]interface{}, error) {
	bytes, err := json.Marshal(c)

	if err != nil {
		return nil, err
	}

	var resultMap map[string]interface{}
	if err := json.Unmarshal(bytes, &resultMap); err != nil {
		return nil, err
	}

	return resultMap, nil
}

//토큰값을 가져오는 것
func JWT() middleware.JWTConfig {
	c := middleware.DefaultJWTConfig
	c.ContextKey = config.Config.Jwt.ContextKey
	c.SigningKey = []byte(config.Config.Jwt.JwtSecret)
	c.Skipper = func(c echo.Context) bool {
		if c.Request().URL.Path == "/" || strings.Contains(c.Request().URL.Path, "/auth/") {
			return true
		}
		return false
	}

	return c
}

//세션 셋팅
func setSession() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !(c.Request().URL.Path == "/" || strings.Contains(c.Request().URL.Path, "/auth/")) {
				requestID := fmt.Sprintf("%v", c.Response().Header().Get(echo.HeaderXRequestID))
				claims := common.UserClaim{}
				rawToken := ""
				token := c.Get("user").(*jwt.Token)
				rawClaims := token.Claims.(jwt.MapClaims)
				rawToken = strings.Split(token.Raw, " ")[0]
				role := fmt.Sprintf("%s", rawClaims["roles"])
				if role == "MEMBER" {
					rawClaims["roles"] = []string{role}
				}
				if err := common.Map2Struct(rawClaims, &claims); err != nil {
					return err
				}
				c.Set(config.ContextUserClaimKey, &common.UserClaim{
					RequestID: requestID,
					ID:        claims.ID,
					Roles:     claims.Roles,
					Name:      claims.Name,
					Datetime:  time.Now().Format(common.DateLayout19),
					Token:     rawToken,
				})
				if claims.ID == 0 {
					return error2.StatusUnauthorized("사용자 자격증명 오류입니다.")
				}
			} else {
				c.Set(config.ContextUserClaimKey, common.UserClaim{})
			}
			return next(c)
		}
	}
}
