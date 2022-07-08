package middleware

import (
	"encoding/json"
	"errors"
	"github.com/YooGenie/daily-work-log-service/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var InvalidAccessToken = errors.New("invalid access token")
var AccessTokenExpired = errors.New("access token expired")

type JwtAuthentication struct {
}

//토큰 발급할 내용을 구조체 한다
type UserClaim struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Roles string `json:"roles"`
}

type JwtToken struct {
	AccessToken  string
	RefreshToken string
}

func (JwtAuthentication) GenerateJwtToken(claim UserClaim) (JwtToken, error) {
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

func (c UserClaim) ConvertMap() (map[string]interface{}, error) {
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
