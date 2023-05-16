package helper

import (
	"strconv"

	t "gitlab.com/azizbekdev-blog/go-monolithic/api/tokens"
	"gitlab.com/azizbekdev-blog/go-monolithic/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func GetClaims(cfg config.Config, c *gin.Context) (*t.CustomClaims, error) {
	var (
		claims = t.CustomClaims{}
	)
	strToken := c.GetHeader("Authorization")

	token, err := jwt.Parse(strToken, func(t *jwt.Token) (interface{}, error) { return []byte(cfg.SignInKey), nil })

	if err != nil {
		return nil, err
	}
	rawClaims := token.Claims.(jwt.MapClaims)

	claims.Sub = rawClaims["sub"].(string)
	claims.Exp = rawClaims["exp"].(float64)
	aud := cast.ToStringSlice(rawClaims["aud"])
	claims.Aud = aud
	claims.Role = rawClaims["role"].(string)
	claims.Sub = rawClaims["sub"].(string)
	claims.Token = token
	return &claims, nil
}

func ParseLimitQueryParam(c *gin.Context) (int, error) {
	return strconv.Atoi(c.DefaultQuery("limit", "10"))
}

func ParsePageQueryParam(c *gin.Context) (int, error) {
	return strconv.Atoi(c.DefaultQuery("page", "1"))
}
